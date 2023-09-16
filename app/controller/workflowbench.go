package controller

import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vpage"
	"github.com/kyleu/loadtoad/views/vworkflow"
	"github.com/valyala/fasthttp"
	"maps"
)

var benchArgs = cutil.Args{
	{Key: "concurrency", Title: "Concurrent Runners", Description: "The number of workflow runners that will execute at once", Default: "1"},
	{Key: "count", Title: "Test Count (per runner)", Description: "The number of workflow runs that will execute for each runner", Default: "1"},
}

func WorkflowStartBench(rc *fasthttp.RequestCtx) {
	Act("workflow.start.bench", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(benchArgs, w, rc)
			if argRes.HasMissing() {
				u := w.WebPath() + "/bench"
				ps.Data = argRes
				return Render(rc, as, &vpage.Args{URL: u, Directions: "Choose your benchmark options", ArgRes: argRes}, ps, "workflow", w.ID, "bench")
			}
			varsStr := argRes.Values["variables"]
			err = util.FromJSON([]byte(varsStr), &w.Variables)
			if err != nil {
				return "", err
			}
			repls = maps.Clone(argRes.Values)
			delete(repls, "variables")
		}

		ents, err := as.Services.LoadToad.LoadEntries(w.Tests...)
		if err != nil {
			return "", err
		}
		ents = ents.WithReplacementsMap(repls, w.Variables)

		ps.Title = "Workflow " + w.Title()
		ps.Data = w
		channel := "run-" + util.RandomString(16)
		p := w.WebPath() + "/connect"
		page := &vworkflow.Start{Workflow: w, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(rc, as, page, ps, "workflow", w.ID, "bench")
	})
}
