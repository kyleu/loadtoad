package controller

import (
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
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
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		ents, err := as.Services.LoadToad.LoadEntries(w.Tests...)
		if err != nil {
			return "", err
		}
		ps.Title = "Workflows"
		ps.Data = w
		return Render(rc, as, &vworkflow.Detail{Workflow: w, Entries: ents}, ps, "workflow", w.ID)
	})
}

func WorkflowStart(rc *fasthttp.RequestCtx) {
	Act("workflow.start", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		ents, err := as.Services.LoadToad.LoadEntries(w.Tests...)
		if err != nil {
			return "", err
		}
		ps.Title = "Workflow " + w.Title()
		ps.Data = w
		channel := "run-" + util.RandomString(16)
		return Render(rc, as, &vworkflow.Start{Workflow: w, Entries: ents, Channel: channel}, ps, "workflow", w.ID, "run")
	})
}

func WorkflowRunSync(rc *fasthttp.RequestCtx) {
	Act("workflow.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		logF := func(i int, s string) {
			ps.Logger.Infof("[%d] %s", i, s)
		}
		ret, err := as.Services.LoadToad.Run(w, logF, nil, nil)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load the toad")
		}
		ps.Title = "Workflow " + w.Title()
		ps.Data = ret
		return Render(rc, as, &vworkflow.Results{Workflow: w, Results: ret}, ps, "workflow", w.ID, "run")
	})
}

func loadWorkflow(as *app.State, rc *fasthttp.RequestCtx) (*loadtoad.Workflow, error) {
	key, err := cutil.RCRequiredString(rc, "key", true)
	if err != nil {
		return nil, err
	}
	return as.Services.LoadToad.LoadWorkflow(key)
}
