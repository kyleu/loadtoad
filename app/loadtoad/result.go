package loadtoad

import (
	"fmt"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/samber/lo"
)

type WorkflowResult struct {
	ID       string        `json:"id"`
	Domain   string        `json:"domain,omitempty"`
	Entry    *har.Entry    `json:"entry,omitempty"`
	Duration int           `json:"duration,omitempty"`
	Logs     []string      `json:"logs,omitempty"`
	Response *har.Response `json:"response,omitempty"`
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
		return x.Duration
	})
}
