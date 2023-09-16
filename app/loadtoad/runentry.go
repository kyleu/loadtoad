package loadtoad

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"slices"

	"github.com/dop251/goja"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

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

func updatedEntry(e *har.Entry, vms map[string]*goja.Runtime) (*har.Entry, error) {
	e = e.Clone()
	for _, vm := range vms {
		err := scriptUpdateEntry(vm, e)
		if err != nil {
			return nil, err
		}
	}
	return e, nil
}
