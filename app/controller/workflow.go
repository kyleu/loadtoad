package controller

import (
	"github.com/kyleu/loadtoad/views/vworkflow"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
)

func WorkflowList(rc *fasthttp.RequestCtx) {
	Act("workflow.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.LoadToad.ListWorkflows(ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return Render(rc, as, &vworkflow.List{Workflows: ret}, ps, "workflow")
	})
}

func WorkflowDetail(rc *fasthttp.RequestCtx) {
	Act("workflow.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.LoadWorkflow(key)
		if err != nil {
			return "", err
		}
		ps.Title = "Workflows"
		ps.Data = ret
		return Render(rc, as, &vworkflow.Detail{Workflow: ret}, ps, "workflow", key)
	})
}

func WorkflowRun(rc *fasthttp.RequestCtx) {
	Act("workflow.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		w, err := as.Services.LoadToad.LoadWorkflow(key)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.Run(w)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load the toad")
		}
		ps.Title = "Workflow " + w.Title()
		ps.Data = ret
		return Render(rc, as, &vworkflow.Results{Workflow: w, Results: ret}, ps, "workflow", key, "run")
	})
}
