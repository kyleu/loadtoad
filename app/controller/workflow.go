package controller

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"strings"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
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

func WorkflowNew(rc *fasthttp.RequestCtx) {
	Act("workflow.new", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w := &loadtoad.Workflow{Tests: har.Selectors{}, Replacements: map[string]string{}, Variables: util.ValueMap{}, Scripts: []string{}}
		arcs := as.Services.Har.List(ps.Logger)
		ps.Title = "New Workflow"
		ps.Data = w
		return Render(rc, as, &vworkflow.Form{Workflow: w, Archives: arcs}, ps, "workflow", "New")
	})
}

func WorkflowCreate(rc *fasthttp.RequestCtx) {
	Act("workflow.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w := &loadtoad.Workflow{}
		err := workflowFromForm(w, rc)
		if err != nil {
			return "", err
		}
		if w.Name == "" {
			return "", errors.New("must provide name")
		}
		if w.ID == "" {
			w.ID = w.Name
		}
		err = as.Services.LoadToad.SaveWorkflow(w)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow created", w.WebPath(), rc, ps)
	})
}

func WorkflowForm(rc *fasthttp.RequestCtx) {
	Act("workflow.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		arcs := as.Services.Har.List(ps.Logger)
		scripts := as.Services.Script.ListScripts(ps.Logger)
		ps.Title = "Edit [" + w.ID + "]"
		ps.Data = w
		return Render(rc, as, &vworkflow.Form{Workflow: w, Archives: arcs, Scripts: scripts}, ps, "workflow", w.ID, "Edit")
	})
}

func WorkflowSave(rc *fasthttp.RequestCtx) {
	Act("workflow.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		err = workflowFromForm(w, rc)
		if err != nil {
			return "", err
		}
		err = as.Services.LoadToad.SaveWorkflow(w)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow saved", w.WebPath(), rc, ps)
	})
}

func WorkflowDelete(rc *fasthttp.RequestCtx) {
	Act("workflow.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		err = as.Services.LoadToad.DeleteWorkflow(w.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow deleted", "/workflow", rc, ps)
	})
}

func collectArgs(args cutil.Args, w *loadtoad.Workflow, rc *fasthttp.RequestCtx) (cutil.Args, *cutil.ArgResults) {
	for k, v := range w.Replacements {
		if strings.Contains(v, "||") {
			choices := util.StringSplitAndTrim(v, "||")
			args = append(args, &cutil.Arg{Key: k, Title: k, Type: "string", Default: "", Choices: choices})
		} else {
			args = append(args, &cutil.Arg{Key: k, Title: k, Type: "string", Default: v})
		}
	}
	args = append(args, &cutil.Arg{Key: "variables", Title: "Other Variables", Type: "textarea", Default: util.ToJSON(w.Variables)})
	return args, cutil.CollectArgs(rc, args)
}
