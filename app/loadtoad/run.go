package loadtoad

import (
	"crypto/tls"
	"fmt"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"net/http"
	"net/http/cookiejar"
)

func (s *Service) LoadEntries(keys ...*har.Selector) (har.Entries, error) {
	var ret har.Entries
	cache := map[string]*har.Log{}
	for _, k := range keys {
		if k.Har == "" {
			return nil, errors.New("each entry must specify an archive")
		}
		h, ok := cache[k.Har]
		if !ok {
			var err error
			h, err = s.LoadHar(k.Har)
			if err != nil {
				return nil, err
			}
			cache[k.Har] = h
		}
		ents, err := h.Entries.Find(k)
		if err != nil {
			return nil, errors.Wrapf(err, "no entries found in [%s] with selector [%s]", h, util.ToJSONCompact(k))
		}
		ents = lo.Filter(ents, func(e *har.Entry, _ int) bool {
			return e.Response != nil && e.Response.Status != 0
		})
		ret = append(ret, ents...)
	}
	return ret, nil
}

func (s *Service) Run(w *Workflow, logF func(i int, s string), errF func(i int, e error), okF func(i int, result *WorkflowResult)) (WorkflowResults, error) {
	var ret WorkflowResults
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	entries, err := s.LoadEntries(w.Tests...)
	if err != nil {
		return nil, err
	}
	var hot []string
	for i, e := range entries {
		wr, err := s.RunEntry(w.ID, i, e, client, jar, hot, logF)
		if err != nil {
			if errF == nil {
				return nil, err
			} else {
				errF(i, err)
			}
		}
		if okF != nil {
			okF(i, wr)
		}
		if !slices.Contains(hot, wr.Domain) {
			hot = append(hot, wr.Domain)
		}
		ret = append(ret, wr)
	}
	return ret, nil
}

func (s *Service) RunEntry(
	wf string, idx int, e *har.Entry, cl http.Client, jar *cookiejar.Jar, hot []string, logF func(int, string),
) (*WorkflowResult, error) {
	id := fmt.Sprintf("%s-%d", wf, idx)
	ret := &WorkflowResult{ID: id, Domain: e.Request.URL, Entry: e.Cleaned()}
	u := e.Request.GetURL()
	if u != nil {
		ret.Domain = u.Host
		if !slices.Contains(hot, u.Host) {
			root := u.Scheme + "://" + u.Host
			req, err := http.NewRequest("GET", root, nil)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Host", u.Host)
			t := util.TimerStart()
			_, _ = cl.Do(req)
			if logF != nil {
				logF(idx, fmt.Sprintf("preconnected to [%s] in [%s]", u.Host, t.EndString()))
			}
		}
	}
	req, err := e.ToRequest(false)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(req.URL, req.Cookies())
	t := util.TimerStart()
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	ret.Duration = t.End()
	ret.Response = har.ResponseFromHTTP(resp)

	if resp != nil {
		_ = resp.Body.Close()
	}
	return ret, nil
}
