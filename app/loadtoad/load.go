package loadtoad

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dop251/goja"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/lib/scripting"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) LoadScripts(scripts []string, logger util.Logger) (map[string]*goja.Runtime, error) {
	vms := make(map[string]*goja.Runtime, len(scripts))
	for _, x := range lo.Uniq(scripts) {
		src, err := s.Script.LoadScript(x, logger)
		if err != nil {
			return nil, err
		}
		_, vm, err := scripting.LoadVM(x, src, logger)
		if err != nil {
			return nil, err
		}
		vms[x] = vm
	}
	return vms, nil
}

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
