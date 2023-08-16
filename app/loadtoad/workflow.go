package loadtoad

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
)

type Workflow struct {
	ID           string            `json:"-"`
	Name         string            `json:"name,omitempty"`
	Tests        har.Selectors     `json:"tests,omitempty"`
	Replacements map[string]string `json:"replacements,omitempty"`
	Variables    util.ValueMap     `json:"variables,omitempty"`
}

func (w *Workflow) Title() string {
	if w.Name == "" {
		return w.ID
	}
	return w.Name
}

func (w *Workflow) WebPath() string {
	return "/workflow/" + url.QueryEscape(w.ID)
}

type Workflows []*Workflow

func NewWorkflow(id string, tests har.Selectors) *Workflow {
	return &Workflow{ID: id, Tests: tests}
}

func (s *Service) SaveWorkflow(w *Workflow) error {
	return s.FS.WriteFile(fmt.Sprintf("workflow/%s.json", w.ID), util.ToJSONBytes(w, true), filesystem.DefaultMode, true)
}

func (s *Service) DeleteWorkflow(id string, logger util.Logger) error {
	return s.FS.Remove(fmt.Sprintf("workflow/%s.json", id), logger)
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
	return ret, nil
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
