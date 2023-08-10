package cmenu

import (
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/samber/lo"
)

func loadtoadMenu(s *loadtoad.Service, logger util.Logger) menu.Items {
	w, _ := s.ListWorkflows(logger)
	workflowKids := lo.Map(w, func(n *loadtoad.Workflow, _ int) *menu.Item {
		return &menu.Item{Key: n.ID, Title: n.Title(), Icon: "sitemap", Route: "/workflow/" + n.ID, Children: menu.Items{
			{Key: "run", Title: "Run", Icon: "play", Route: "/workflow/" + n.ID + "/run"},
		}}
	})
	workflow := &menu.Item{Key: "workflow", Title: "Workflows", Description: "Workflows for Load Toad!", Icon: "sitemap", Route: "/workflow", Children: workflowKids}

	harKids := lo.Map(s.ListHars(logger), func(n string, _ int) *menu.Item {
		return &menu.Item{Key: n, Title: n, Icon: "book", Route: "/har/" + n}
	})
	har := &menu.Item{Key: "har", Title: "Archives", Description: "HTTP Archive files", Icon: "book", Route: "/har", Children: harKids}

	return menu.Items{workflow, menu.Separator, har}
}
