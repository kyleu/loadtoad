package cmenu

import (
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

func loadtoadMenu(s *loadtoad.Service, logger util.Logger) *menu.Item {
	w, _ := s.ListWorkflows(logger)
	workflowKids := lo.Map(w, func(n *loadtoad.Workflow, _ int) *menu.Item {
		return &menu.Item{Key: n.ID, Title: n.Title(), Icon: "sitemap", Route: n.WebPath(), Children: menu.Items{
			{Key: "run", Title: "Run", Icon: "file-code", Route: n.WebPath() + "/run"},
			{Key: "bench", Title: "Benchmark", Icon: "laptop-code", Route: n.WebPath() + "/benchmark"},
		}}
	})
	wfDesc := "Workflows for Load Toad!"
	return &menu.Item{Key: "workflow", Title: "Workflows", Description: wfDesc, Icon: "sitemap", Route: "/workflow", Children: workflowKids}
}
