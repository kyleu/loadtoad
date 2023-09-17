package controller

import (
	"context"
	"maps"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

func WorkflowConnectBench(rc *fasthttp.RequestCtx) {
	Act("workflow.connect.bench", rc, workflowConnect(rc, "bench"))
}

func WorkflowConnectRun(rc *fasthttp.RequestCtx) {
	Act("workflow.connect.run", rc, workflowConnect(rc, "run"))
}

func workflowConnect(rc *fasthttp.RequestCtx, key string) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		repls := maps.Clone(w.Replacements)
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			argRes := collectArgs(key, w, rc)
			if argRes.HasMissing() {
				return "", errors.Errorf("missing arguments [%s]", strings.Join(argRes.Missing, ", "))
			}
			for k, v := range argRes.Values {
				switch k {
				case "count", "concurrency":
				// noop
				case "variables":
					others := util.ValueMap{}
					_ = util.FromJSON([]byte(v), &others)
					w.Variables = others.Merge(w.Variables)
				default:
					repls[k] = v
				}
			}
		}
		return socketConnect(ps.Context, key, w, repls, rc, as, ps)
	}
}

func socketConnect(
	ctx context.Context, key string, w *loadtoad.Workflow, repls map[string]string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState,
) (string, error) {
	send, logF, errF, okF, err := wireSocketFuncs(rc, as, ps)
	if err != nil {
		return "", err
	}
	go func() {
		var final loadtoad.WorkflowResults
		var e error
		if key == "bench" {
			final, e = as.Services.LoadToad.BenchWorkflow(ctx, w, repls, ps.Logger, logF, errF, okF)
		} else {
			final, e = as.Services.LoadToad.RunWorkflow(ctx, w, repls, ps.Logger, logF, errF, okF)
		}
		if e != nil {
			errF(-1, e)
		}
		msg := map[string]any{"message": util.MicrosToMillis(final.Duration()), "status": "Success"}
		ps.Logger.Infof("[COMPLETE] %s", w.ID)
		send("complete", &WorkflowMessage{Idx: -1, Ctx: msg})
	}()
	return "", nil
}
