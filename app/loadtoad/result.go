package loadtoad

import (
	"fmt"

	"github.com/samber/lo"

	har2 "github.com/kyleu/loadtoad/app/lib/har"
)

type WorkflowResult struct {
	ID       string            `json:"id"`
	Domain   string            `json:"domain,omitempty"`
	Entry    *har2.Entry       `json:"entry,omitempty"`
	Timing   *har2.PageTimings `json:"duration,omitempty"`
	Logs     []string          `json:"logs,omitempty"`
	Response *har2.Response    `json:"response,omitempty"`
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
