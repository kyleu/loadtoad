package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func WorkflowList(w http.ResponseWriter, r *http.Request) {
	Act("workflow.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.LoadToad.ListWorkflows(ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return Render(r, as, &vworkflow.List{Workflows: ret}, ps, "workflow")
	})
}

func WorkflowDetail(w http.ResponseWriter, r *http.Request) {
	Act("workflow.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, err := loadWorkflow(as, r)
		if err != nil {
			return "", err
		}
		ents, err := as.Services.LoadToad.LoadEntries(wf.Tests...)
		if err != nil {
			return "", err
		}
		ps.Title = "Workflows"
		ps.Data = wf
		return Render(r, as, &vworkflow.Detail{Workflow: wf, Entries: ents}, ps, "workflow", wf.ID)
	})
}

func WorkflowNew(w http.ResponseWriter, r *http.Request) {
	Act("workflow.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf := &loadtoad.Workflow{Tests: har.Selectors{}, Replacements: map[string]string{}, Variables: util.ValueMap{}, Scripts: []string{}}
		arcs := as.Services.Har.List(ps.Logger)
		ps.Title = "New Workflow"
		ps.Data = wf
		return Render(r, as, &vworkflow.Form{Workflow: wf, Archives: arcs}, ps, "workflow", "New")
	})
}

func WorkflowCreate(w http.ResponseWriter, r *http.Request) {
	Act("workflow.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf := &loadtoad.Workflow{}
		err := workflowFromForm(wf, r, ps.RequestBody)
		if err != nil {
			return "", err
		}
		if wf.Name == "" {
			return "", errors.New("must provide name")
		}
		if wf.ID == "" {
			wf.ID = wf.Name
		}
		err = as.Services.LoadToad.SaveWorkflow(wf)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow created", wf.WebPath(), ps)
	})
}

func WorkflowForm(w http.ResponseWriter, r *http.Request) {
	Act("workflow.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, err := loadWorkflow(as, r)
		if err != nil {
			return "", err
		}
		arcs := as.Services.Har.List(ps.Logger)
		scripts := as.Services.Script.ListScripts(ps.Logger)
		ps.Title = "Edit [" + wf.ID + "]"
		ps.Data = wf
		return Render(r, as, &vworkflow.Form{Workflow: wf, Archives: arcs, Scripts: scripts}, ps, "workflow", wf.ID, "Edit")
	})
}

func WorkflowSave(w http.ResponseWriter, r *http.Request) {
	Act("workflow.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, err := loadWorkflow(as, r)
		if err != nil {
			return "", err
		}
		err = workflowFromForm(wf, r, ps.RequestBody)
		if err != nil {
			return "", err
		}
		err = as.Services.LoadToad.SaveWorkflow(wf)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow saved", wf.WebPath(), ps)
	})
}

func WorkflowDelete(w http.ResponseWriter, r *http.Request) {
	Act("workflow.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, err := loadWorkflow(as, r)
		if err != nil {
			return "", err
		}
		err = as.Services.LoadToad.DeleteWorkflow(wf.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return FlashAndRedir(true, "Workflow deleted", "/workflow", ps)
	})
}
