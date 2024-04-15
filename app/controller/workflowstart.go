package controller

import (
	"fmt"
	"maps"
	"net/http"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vpage"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func WorkflowStartRun(w http.ResponseWriter, r *http.Request) {
	Act("workflow.start.run", w, r, workflowStart(w, r, "run"))
}

func WorkflowStartBench(w http.ResponseWriter, r *http.Request) {
	Act("workflow.start.bench", w, r, workflowStart(w, r, "bench"))
}

func workflowStart(w http.ResponseWriter, r *http.Request, key string) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, repls, argRes, err := loadWorkflowWithRepls(key, as, r)
		if err != nil {
			return "", err
		}
		if argRes != nil && argRes.HasMissing() {
			u := fmt.Sprintf("%s/%s", wf.WebPath(), key)
			ps.Data = argRes
			return Render(r, as, &vpage.Args{URL: u, Directions: "Choose your benchmark options", ArgRes: argRes}, ps, "workflow", wf.ID, "bench")
		}

		ents, err := as.Services.LoadToad.LoadEntries(wf.Tests...)
		if err != nil {
			return "", err
		}
		ents = ents.WithReplacementsMap(repls, wf.Variables)

		ps.Title = fmt.Sprintf("%s [%s]", wf.Title(), key)
		ps.Data = w
		channel := fmt.Sprintf("%s-%s", key, util.RandomString(16))
		p := fmt.Sprintf("%s/%s/connect", wf.WebPath(), key)
		page := &vworkflow.Start{Workflow: wf, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(r, as, page, ps, "workflow", wf.ID, key)
	}
}

func loadWorkflowWithRepls(key string, as *app.State, r *http.Request) (*loadtoad.Workflow, map[string]string, *cutil.ArgResults, error) {
	wf, err := loadWorkflow(as, r)
	if err != nil {
		return nil, nil, nil, err
	}
	repls := maps.Clone(wf.Replacements)
	if r.URL.Query().Get("ok") != util.BoolTrue {
		argRes := collectArgs(key, wf, r)
		if argRes.HasMissing() {
			return wf, nil, argRes, nil
		}
		varsStr := argRes.Values.GetStringOpt("variables")
		err = util.FromJSON([]byte(varsStr), &wf.Variables)
		if err != nil {
			return wf, repls, argRes, err
		}
		repls = maps.Clone(argRes.Values.ToStringMap())
		delete(repls, "variables")
	}
	return wf, repls, nil, err
}
