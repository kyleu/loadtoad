package controller

import (
	"fmt"
	"maps"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vpage"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func WorkflowStartRun(rc *fasthttp.RequestCtx) {
	Act("workflow.start.run", rc, workflowStart(rc, "run"))
}

func WorkflowStartBench(rc *fasthttp.RequestCtx) {
	Act("workflow.start.bench", rc, workflowStart(rc, "bench"))
}

func workflowStart(rc *fasthttp.RequestCtx, key string) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		w, repls, argRes, err := loadWorkflowWithRepls(key, as, rc)
		if err != nil {
			return "", err
		}
		if argRes != nil && argRes.HasMissing() {
			u := fmt.Sprintf("%s/%s", w.WebPath(), key)
			ps.Data = argRes
			return Render(rc, as, &vpage.Args{URL: u, Directions: "Choose your benchmark options", ArgRes: argRes}, ps, "workflow", w.ID, "bench")
		}

		ents, err := as.Services.LoadToad.LoadEntries(w.Tests...)
		if err != nil {
			return "", err
		}
		ents = ents.WithReplacementsMap(repls, w.Variables)

		ps.Title = fmt.Sprintf("%s [%s]", w.Title(), key)
		ps.Data = w
		channel := fmt.Sprintf("%s-%s", key, util.RandomString(16))
		p := fmt.Sprintf("%s/%s/connect", w.WebPath(), key)
		page := &vworkflow.Start{Workflow: w, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(rc, as, page, ps, "workflow", w.ID, key)
	}
}

func loadWorkflowWithRepls(key string, as *app.State, rc *fasthttp.RequestCtx) (*loadtoad.Workflow, map[string]string, *cutil.ArgResults, error) {
	w, err := loadWorkflow(as, rc)
	if err != nil {
		return nil, nil, nil, err
	}
	repls := maps.Clone(w.Replacements)
	if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
		argRes := collectArgs(key, w, rc)
		if argRes.HasMissing() {
			return w, nil, argRes, nil
		}
		varsStr := argRes.Values["variables"]
		err = util.FromJSON([]byte(varsStr), &w.Variables)
		if err != nil {
			return w, repls, argRes, err
		}
		repls = maps.Clone(argRes.Values)
		delete(repls, "variables")
	}
	return w, repls, nil, err
}
