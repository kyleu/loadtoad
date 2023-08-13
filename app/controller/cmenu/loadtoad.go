package cmenu

import (
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
)

func loadtoadMenu(s *loadtoad.Service, logger util.Logger) menu.Items {
	w, _ := s.ListWorkflows(logger)
	workflowKids := lo.Map(w, func(n *loadtoad.Workflow, _ int) *menu.Item {
		return &menu.Item{Key: n.ID, Title: n.Title(), Icon: "sitemap", Route: "/workflow/" + n.ID, Children: menu.Items{
			{Key: "run", Title: "Run", Icon: "file-code", Route: "/workflow/" + n.ID + "/run"},
			{Key: "bench", Title: "Benchmark", Icon: "laptop-code", Route: "/workflow/" + n.ID + "/benchmark"},
		}}
	})
	wfDesc := "Workflows for Load Toad!"
	workflow := &menu.Item{Key: "workflow", Title: "Workflows", Description: wfDesc, Icon: "sitemap", Route: "/workflow", Children: workflowKids}

	harKids := lo.Map(s.ListHars(logger), func(n string, _ int) *menu.Item {
		n = strings.TrimSuffix(n, har.Ext)
		return &menu.Item{Key: n, Title: n, Icon: "book", Route: "/har/" + n}
	})
	harItem := &menu.Item{Key: "har", Title: "Archives", Description: "HTTP Archive files", Icon: "book", Route: "/har", Children: harKids}

	return menu.Items{workflow, menu.Separator, harItem}
}
