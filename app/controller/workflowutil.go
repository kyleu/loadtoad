package controller

import (
	"github.com/kyleu/loadtoad/app/util"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
)

func loadWorkflow(as *app.State, rc *fasthttp.RequestCtx) (*loadtoad.Workflow, error) {
	key, err := cutil.RCRequiredString(rc, "key", true)
	if err != nil {
		return nil, err
	}
	return as.Services.LoadToad.LoadWorkflow(key)
}

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
