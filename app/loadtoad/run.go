package loadtoad

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"slices"

	har2 "github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) Run(
	ctx context.Context, w *Workflow, repls map[string]string, logger util.Logger,
	logF func(i int, s string), errF func(i int, e error), okF func(i int, result *WorkflowResult),
) (WorkflowResults, error) {
	var ret WorkflowResults
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}} //nolint:gosec
	_, entries, err := s.LoadEntries(repls, w.Variables, w.Scripts, logger, w.Tests...)
	if err != nil {
		return nil, err
	}
	var hot []string
	for i, e := range entries {
		cl := e.Clone()
		wr, err := s.RunEntry(ctx, w.ID, i, cl, client, jar, hot, logF)
		if err == nil {
			if okF != nil {
				okF(i, wr)
			}
			if !slices.Contains(hot, wr.Domain) {
				hot = append(hot, wr.Domain)
			}
			ret = append(ret, wr)
		} else {
			if errF == nil {
				return nil, err
			}
			errF(i, err)
		}
	}
	return ret, nil
}

func (s *Service) RunEntry(
	ctx context.Context, wf string, idx int, e *har2.Entry, cl http.Client, jar *cookiejar.Jar, hot []string, logF func(int, string),
) (*WorkflowResult, error) {
	id := fmt.Sprintf("%s-%d", wf, idx)
	ret := &WorkflowResult{ID: id, Domain: e.Request.URL, Entry: e.Cleaned(), Timing: &har2.PageTimings{}}
	u := e.Request.GetURL()
	if u != nil {
		ret.Domain = u.Host
		if !slices.Contains(hot, u.Host) {
			if err := preload(ctx, u.Scheme, u.Host, e.Request.Headers, idx, cl, logF); err != nil {
				return nil, err
			}
		}
	}
	req, err := e.ToRequest(ctx, false)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(req.URL, req.Cookies())
	req = wireReq(req, ret.Timing)

	t := util.TimerStart()
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	ret.Timing.Total = t.End()
	ret.Response = har2.ResponseFromHTTP(resp)

	if resp != nil {
		_ = resp.Body.Close()
	}
	return ret, nil
}
