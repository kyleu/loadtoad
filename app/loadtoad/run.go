package loadtoad

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"slices"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) Run(
	ctx context.Context, w *Workflow, repls map[string]string, logger util.Logger,
	logF func(i int, s string), errF func(i int, e error), okF func(i int, result *WorkflowResult),
) (WorkflowResults, error) {
	var ret WorkflowResults
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}} //nolint:gosec
	entries, err := s.LoadEntries(w.Tests...)
	if err != nil {
		return nil, err
	}
	vms, err := s.LoadScripts(w.Scripts, logger)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		for _, vm := range vms {
			newRepls, err2 := scriptExtractReplacements(vm, e)
			if err2 != nil {
				return nil, err2
			}
			if len(newRepls) > 0 {
				for k, v := range newRepls {
					repls[k] = v
				}
			}
		}
	}
	vars := w.Variables.Clone()

	var hot []string
	for i, e := range entries {
		e = e.Clone()
		for _, vm := range vms {
			err = scriptUpdateEntry(vm, e)
			if err != nil {
				return nil, err
			}
		}
		var wr *WorkflowResult
		wr, err = s.RunEntry(ctx, w.ID, i, e, client, jar, hot, repls, vars, logF)
		if err == nil {
			for _, vm := range vms {
				newVars, err2 := scriptExtractVariables(vm, wr)
				if err2 != nil {
					return nil, err2
				}
				if len(newVars) > 0 {
					title := util.StringPlural(len(newVars), "variable")
					keys := strings.Join(lo.Keys(newVars), ", ")
					logF(i, fmt.Sprintf("observed [%s] (%s)", title, keys))
					vars = vars.Merge(newVars)
				}
			}
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
	ctx context.Context, wf string, idx int, e *har.Entry, cl http.Client, jar *cookiejar.Jar,
	hot []string, repls map[string]string, vars util.ValueMap, logF func(int, string),
) (*WorkflowResult, error) {
	e = e.WithReplacementsMap(repls, vars)
	id := fmt.Sprintf("%s-%d", wf, idx)
	ret := &WorkflowResult{ID: id, Domain: e.Request.URL, Entry: e.Cleaned(), Replacements: repls, Variables: vars.Clone(), Timing: &har.PageTimings{}}
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
	ret.Response, err = har.ResponseFromHTTP(resp)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		_ = resp.Body.Close()
	}
	return ret, nil
}
