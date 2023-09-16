package loadtoad

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/filter"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/lib/scripting"
	"github.com/kyleu/loadtoad/app/lib/search/result"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/util"
)

type Service struct {
	FS     filesystem.FileLoader
	Har    *har.Service
	Socket *websocket.Service
	Script *scripting.Service
}

func NewService(fs filesystem.FileLoader, harSvc *har.Service, ws *websocket.Service, script *scripting.Service) *Service {
	ret := &Service{FS: fs, Har: harSvc, Socket: ws, Script: script}
	return ret
}

func (s *Service) ListWorkflows(logger util.Logger) (Workflows, error) {
	names := s.FS.ListExtension("workflow", "json", nil, true, logger)
	ret := make(Workflows, 0, len(names))
	for _, n := range names {
		w, err := s.LoadWorkflow(n)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to load workflow [%s]", n)
		}
		ret = append(ret, w)
	}
	return ret, nil
}

func (s *Service) LoadWorkflow(fn string) (*Workflow, error) {
	if !strings.HasSuffix(fn, ".json") {
		fn += ".json"
	}
	if !strings.Contains(fn, "workflow") {
		fn = path.Join("workflow", fn)
	}
	if !s.FS.Exists(fn) {
		return nil, errors.Errorf("missing workflow file [%s]", fn)
	}
	b, err := s.FS.ReadFile(fn)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading workflow file [%s]", fn)
	}
	ret := &Workflow{}
	err = util.FromJSON(b, ret)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding workflow file [%s]", fn)
	}
	if ret.ID == "" {
		ret.ID = strings.TrimPrefix(strings.TrimSuffix(fn, ".json"), "workflow/")
	}
	for k, v := range ret.Replacements {
		if v == "" {
			ret.Replacements[k] = k
		}
	}
	if ret.Replacements == nil {
		ret.Replacements = map[string]string{}
	}
	if ret.Variables == nil {
		ret.Variables = util.ValueMap{}
	}
	if ret.Scripts == nil {
		ret.Scripts = []string{}
	}
	return ret, nil
}

func (s *Service) SaveWorkflow(w *Workflow) error {
	p := fmt.Sprintf("workflow/%s.json", w.ID)
	content := util.ToJSONBytes(w, true)
	return s.FS.WriteFile(p, content, filesystem.DefaultMode, true)
}

func (s *Service) DeleteWorkflow(id string, logger util.Logger) error {
	return s.FS.Remove(fmt.Sprintf("workflow/%s.json", id), logger)
}

func (s *Service) SearchWorkflows(_ context.Context, _ filter.ParamSet, q string, logger util.Logger) (result.Results, error) {
	wfs, err := s.ListWorkflows(logger)
	if err != nil {
		return nil, err
	}
	f := func(w *Workflow, _ int) (*result.Result, bool) {
		res := result.NewResult("workflow", w.ID, w.WebPath(), w.Title(), "sitemap", w, w, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}
	return lo.FilterMap(wfs, f), nil
}
