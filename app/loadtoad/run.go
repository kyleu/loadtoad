package loadtoad

import (
	"crypto/tls"
	"fmt"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/pkg/errors"
	"net/http"
	"net/http/cookiejar"
)

func (s *Service) Run(w *Workflow) (WorkflowResults, error) {
	cache := map[string]*har.Log{}
	var ret WorkflowResults

	jar, _ := cookiejar.New(nil)

	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	for i, k := range w.RequestKeys {
		h, ok := cache[k.HAR]
		if !ok {
			var err error
			h, err = s.LoadHar(k.HAR)
			if err != nil {
				return nil, err
			}
			cache[k.HAR] = h
		}
		ents := h.Entries
		if k.Page != "" {
			ents = ents.ForPage(k.Page)
		}
		e := ents.Find(k.URL, k.Idx)
		if e == nil {
			return nil, errors.Errorf("no entry found in [%s] with url [%s] or index [%d]", k.Page, k.URL, k.Idx)
		}
		id := fmt.Sprintf("%s-%d", w.ID, i)
		wr, err := s.RunEntry(id, e, client, jar)
		if err != nil {
			return nil, err
		}
		ret = append(ret, wr)
	}
	return ret, nil
}

func (s *Service) RunEntry(id string, e *har.Entry, client http.Client, jar *cookiejar.Jar) (*WorkflowResult, error) {
	ret := &WorkflowResult{ID: id, Entry: e}
	t := util.TimerStart()
	req, err := e.ToRequest(false)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(req.URL, req.Cookies())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	ret.Response = har.ResponseFromHTTP(resp)

	if resp != nil {
		_ = resp.Body.Close()
	}
	ret.Duration = t.End()
	return ret, nil
}
