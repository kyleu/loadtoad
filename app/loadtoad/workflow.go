package loadtoad

import (
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/pkg/errors"
	"path"
	"strings"
)

type Workflow struct {
	ID    string        `json:"-"`
	Name  string        `json:"name,omitempty"`
	Tests har.Selectors `json:"tests,omitempty"`
}

func (w *Workflow) Title() string {
	if w.Name == "" {
		return w.ID
	}
	return w.Name
}

func (w *Workflow) WebPath() string {
	return "/workflow/" + w.ID
}

type Workflows []*Workflow

func NewWorkflow(id string, tests har.Selectors) *Workflow {
	return &Workflow{ID: id, Tests: tests}
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
	return ret, nil
}

func (s *Service) ListWorkflows(logger util.Logger) (Workflows, error) {
	names := s.FS.ListExtension("workflow", "json", nil, true, logger)
	var ret Workflows
	for _, n := range names {
		w, err := s.LoadWorkflow(n)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to load workflow [%s]", n)
		}
		ret = append(ret, w)
	}
	return ret, nil
}
