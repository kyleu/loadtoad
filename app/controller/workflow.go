package controller

import (
	"fmt"
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vpage"
	"github.com/kyleu/loadtoad/views/vworkflow"
	"github.com/valyala/fasthttp"
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
		ents, err := as.Services.LoadToad.LoadEntries(w.Replacements, w.Tests...)
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

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != "true" {
			var args cutil.Args
			for k, v := range w.Replacements {
				args = append(args, &cutil.Arg{Key: k, Title: k, Type: "string", Default: v})
			}
			argRes := cutil.CollectArgs(rc, args)
			if argRes.HasMissing() {
				url := fmt.Sprintf("/workflow/%s/run", w.ID)
				ps.Data = argRes
				return Render(rc, as, &vpage.Args{URL: url, Directions: "Choose your run options", ArgRes: argRes}, ps, "workflow", w.ID, "run")
			}
			repls = argRes.Values
		}

		ents, err := as.Services.LoadToad.LoadEntries(repls, w.Tests...)
		if err != nil {
			return "", err
		}
		ps.Title = "Workflow " + w.Title()
		ps.Data = w
		channel := "run-" + util.RandomString(16)
		return Render(rc, as, &vworkflow.Start{Workflow: w, Entries: ents, Channel: channel, Path: "/workflow/" + w.ID + "/connect"}, ps, "workflow", w.ID, "run")
	})
}

func loadWorkflow(as *app.State, rc *fasthttp.RequestCtx) (*loadtoad.Workflow, error) {
	key, err := cutil.RCRequiredString(rc, "key", true)
	if err != nil {
		return nil, err
	}
	return as.Services.LoadToad.LoadWorkflow(key)
}
