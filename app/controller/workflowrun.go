package controller

import (
	"context"
	"maps"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vpage"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func WorkflowStartRun(rc *fasthttp.RequestCtx) {
	Act("workflow.start.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(nil, w, rc)
			if argRes.HasMissing() {
				u := w.WebPath() + "/run"
				ps.Data = argRes
				return Render(rc, as, &vpage.Args{URL: u, Directions: "Choose your run options", ArgRes: argRes}, ps, "workflow", w.ID, "run")
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
		p := "/workflow/" + url.QueryEscape(w.ID) + "/connect"
		page := &vworkflow.Start{Workflow: w, Entries: ents.Cleaned(), Replacements: repls, Channel: channel, Path: p}
		return Render(rc, as, page, ps, "workflow", w.ID, "run")
	})
}

func WorkflowConnectRun(rc *fasthttp.RequestCtx) {
	Act("workflow.connect.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(nil, w, rc)
			if argRes.HasMissing() {
				return "", errors.Errorf("missing arguments [%s]", strings.Join(argRes.Missing, ", "))
			}
			for k, v := range argRes.Values {
				repls[k] = v
			}
		}
		return socketConnectRun(ps.Context, w, repls, rc, as, ps)
	})
}

func socketConnectRun(
	ctx context.Context, w *loadtoad.Workflow, repls map[string]string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState,
) (string, error) {
	send, logF, errF, okF, err := wireSocketFuncs(rc, as, ps)
	if err != nil {
		return "", err
	}
	go func() {
		final, e := as.Services.LoadToad.RunWorkflow(ctx, w, repls, ps.Logger, logF, errF, okF)
		if e != nil {
			errF(-1, e)
		}
		ps.Logger.Infof("[COMPLETE] %s", w.ID)
		msg := map[string]any{"message": util.MicrosToMillis(final.Duration()), "status": "Success"}
		send("complete", &WorkflowMessage{Idx: -1, Ctx: msg})
	}()
	return "", nil
}
