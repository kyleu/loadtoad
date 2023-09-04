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

func workflowFromForm(w *loadtoad.Workflow, rc *fasthttp.RequestCtx) error {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return err
	}

	w.Name = frm.GetStringOpt("name")

	err = util.FromJSON([]byte(frm.GetStringOpt("tests")), &w.Tests)
	if err != nil {
		return err
	}

	repls := map[string]string{}
	err = util.FromJSON([]byte(frm.GetStringOpt("replacements")), &repls)
	if err != nil {
		return err
	}
	w.Replacements = repls

	vars := util.ValueMap{}
	err = util.FromJSON([]byte(frm.GetStringOpt("variables")), &vars)
	if err != nil {
		return err
	}
	w.Variables = vars

	err = util.FromJSON([]byte(frm.GetStringOpt("scripts")), &w.Scripts)
	if err != nil {
		return err
	}
	return nil
}

func WorkflowStartRun(rc *fasthttp.RequestCtx) {
	Act("workflow.start.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(w, rc)
			if argRes.HasMissing() {
				url := fmt.Sprintf("/workflow/%s/run", w.ID)
				ps.Data = argRes
				return Render(rc, as, &vpage.Args{URL: url, Directions: "Choose your run options", ArgRes: argRes}, ps, "workflow", w.ID, "run")
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
		p := "/workflow/" + w.ID + "/connect"
		page := &vworkflow.Start{Workflow: w, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(rc, as, page, ps, "workflow", w.ID, "run")
	})
}

func WorkflowStartBench(rc *fasthttp.RequestCtx) {
	Act("workflow.start.bench", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(w, rc)
			if argRes.HasMissing() {
				url := fmt.Sprintf("/workflow/%s/run", w.ID)
				ps.Data = argRes
				return Render(rc, as, &vpage.Args{URL: url, Directions: "Choose your benchmark options", ArgRes: argRes}, ps, "workflow", w.ID, "bench")
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
		p := "/workflow/" + w.ID + "/connect"
		page := &vworkflow.Start{Workflow: w, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(rc, as, page, ps, "workflow", w.ID, "run")
	})
}
