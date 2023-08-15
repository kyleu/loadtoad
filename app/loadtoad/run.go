package loadtoad

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) LoadEntries(repls map[string]string, vars util.ValueMap, keys ...*har.Selector) (har.Entries, error) {
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
			return nil, errors.Wrapf(err, "no entries found in [%s] with selector [%s]", k.Har, util.ToJSONCompact(k))
		}
		ents = lo.Filter(ents, func(e *har.Entry, _ int) bool {
			return e.Response != nil && e.Response.Status != 0
		})
		lo.ForEach(ents, func(x *har.Entry, _ int) {
			x.Selector = k
		})
		ret = append(ret, ents...)
	}
	return ret.WithReplacementsMap(repls, vars), nil
}

func (s *Service) Run(
	ctx context.Context, w *Workflow, repls map[string]string, logF func(i int, s string), errF func(i int, e error), okF func(i int, result *WorkflowResult),
) (WorkflowResults, error) {
	var ret WorkflowResults
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}} //nolint:gosec
	entries, err := s.LoadEntries(repls, w.Variables, w.Tests...)
	if err != nil {
		return nil, err
	}
	var hot []string
	for i, e := range entries {
		wr, err := s.RunEntry(ctx, w.ID, i, e, client, jar, hot, logF)
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
	ctx context.Context, wf string, idx int, e *har.Entry, cl http.Client, jar *cookiejar.Jar, hot []string, logF func(int, string),
) (*WorkflowResult, error) {
	id := fmt.Sprintf("%s-%d", wf, idx)
	ret := &WorkflowResult{ID: id, Domain: e.Request.URL, Entry: e.Cleaned(), Timing: &har.PageTimings{}}
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
	ret.Response = har.ResponseFromHTTP(resp)

	if resp != nil {
		_ = resp.Body.Close()
	}
	return ret, nil
}

func preload(ctx context.Context, scheme string, host string, headers har.NVPs, idx int, cl http.Client, logF func(int, string)) error {
	root := scheme + "://" + host
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, root, http.NoBody)
	if err != nil {
		return err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", host)
	for _, h := range headers {
		req.Header.Set(h.Name, h.Value)
	}
	t := util.TimerStart()
	rsp, err := cl.Do(req)
	if err != nil {
		logF(idx, fmt.Sprintf("error preconnecting to [%s]: %v", host, err))
	}
	if rsp != nil && rsp.Body != nil {
		defer func() {
			_ = rsp.Body.Close()
		}()
		_, _ = io.ReadAll(rsp.Body)
	}
	if logF != nil {
		logF(idx, fmt.Sprintf("preconnected to [%s] in [%s]", host, t.EndString()))
	}
	return nil
}

func wireReq(req *http.Request, timing *har.PageTimings) *http.Request {
	start := util.TimerStart()
	var connect, dns, tlsHandshake *util.Timer
	timing.IsMicros = true

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = util.TimerStart() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			timing.DNS = dns.End()
		},

		TLSHandshakeStart: func() { tlsHandshake = util.TimerStart() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			timing.SSL = tlsHandshake.End()
		},

		ConnectStart: func(network, addr string) { connect = util.TimerStart() },
		ConnectDone: func(network, addr string, err error) {
			timing.Connect = connect.End()
		},

		GotFirstResponseByte: func() {
			timing.Receive = start.End()
		},
	}
	return req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
}
