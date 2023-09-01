package loadtoad

import (
	"fmt"

	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

type WorkflowResult struct {
	ID           string            `json:"id"`
	Domain       string            `json:"domain,omitempty"`
	Entry        *har.Entry        `json:"entry,omitempty"`
	Replacements map[string]string `json:"repls,omitempty"`
	Variables    util.ValueMap     `json:"vars,omitempty"`
	Timing       *har.PageTimings  `json:"duration,omitempty"`
	Logs         []string          `json:"logs,omitempty"`
	Response     *har.Response     `json:"response,omitempty"`
}

func (w *WorkflowResult) Title() string {
	return w.Entry.String()
}

func (w *WorkflowResult) AddLog(s string, args ...any) {
	w.Logs = append(w.Logs, fmt.Sprintf(s, args...))
}

type WorkflowResults []*WorkflowResult

func (w WorkflowResults) Duration() int {
	return lo.SumBy(w, func(x *WorkflowResult) int {
		return x.Timing.Total
	})
}
