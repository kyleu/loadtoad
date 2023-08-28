package loadtoad

import (
	"net/url"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

type Workflow struct {
	ID           string            `json:"-"`
	Name         string            `json:"name,omitempty"`
	Tests        har.Selectors     `json:"tests,omitempty"`
	Replacements map[string]string `json:"replacements,omitempty"`
	Variables    util.ValueMap     `json:"variables,omitempty"`
	Scripts      []string          `json:"scripts,omitempty"`
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
