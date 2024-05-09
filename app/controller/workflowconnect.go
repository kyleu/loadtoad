package controller

import (
	"context"
	"maps"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

func WorkflowConnectBench(w http.ResponseWriter, r *http.Request) {
	Act("workflow.connect.bench", w, r, workflowConnect(w, r, "bench"))
}

func WorkflowConnectRun(w http.ResponseWriter, r *http.Request) {
	Act("workflow.connect.run", w, r, workflowConnect(w, r, "run"))
}

func workflowConnect(w http.ResponseWriter, r *http.Request, key string) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		wf, err := loadWorkflow(as, r)
		if err != nil {
			return "", err
		}
		repls := maps.Clone(wf.Replacements)
		if r.URL.Query().Get("ok") != util.BoolTrue {
			argRes := collectArgs(key, wf, r)
			if argRes.HasMissing() {
				return "", errors.Errorf("missing arguments [%s]", strings.Join(argRes.Missing, ", "))
			}
			for k := range argRes.Values {
				v := argRes.Values.GetStringOpt(k)
				switch k {
				case "count", "concurrency":
				// noop
				case "variables":
					others := util.ValueMap{}
					_ = util.FromJSON([]byte(v), &others)
					wf.Variables = others.Merge(wf.Variables)
				default:
					repls[k] = v
				}
			}
		}
		return socketConnect(ps.Context, key, wf, repls, w, r, as, ps)
	}
}

func socketConnect(
	ctx context.Context, key string, wf *loadtoad.Workflow, repls map[string]string, w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState,
) (string, error) {
	id, send, logF, errF, okF, err := wireSocketFuncs(w, r, as, ps)
	if err != nil {
		return "", err
	}
	go func() {
		var final loadtoad.WorkflowResults
		var e error
		if key == "bench" {
			final, e = as.Services.LoadToad.BenchWorkflow(ctx, wf, repls, ps.Logger, logF, errF, okF)
		} else {
			final, e = as.Services.LoadToad.RunWorkflow(ctx, wf, repls, ps.Logger, logF, errF, okF)
		}
		if e != nil {
			errF(-1, e)
		}
		msg := map[string]any{"message": util.MicrosToMillis(final.Duration()), "status": "Success"}
		ps.Logger.Infof("[COMPLETE] %s", wf.ID)
		send("complete", &WorkflowMessage{Idx: -1, Ctx: msg})
	}()
	if err = as.Services.Socket.ReadLoop(context.Background(), id, ps.Logger); err != nil {
		return "", err
	}
	return "", nil
}
