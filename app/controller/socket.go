package controller

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

type WorkflowMessage struct {
	Idx int `json:"idx"`
	Ctx any `json:"ctx,omitempty"`
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

func socketConnectBench(
	ctx context.Context, w *loadtoad.Workflow, repls map[string]string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState,
) (string, error) {
	send, logF, errF, okF, err := wireSocketFuncs(rc, as, ps)
	if err != nil {
		return "", err
	}
	go func() {
		final, e := as.Services.LoadToad.BenchWorkflow(ctx, w, repls, ps.Logger, logF, errF, okF)
		if e != nil {
			errF(-1, e)
		}
		ps.Logger.Infof("[COMPLETE] %s", w.ID)
		msg := map[string]any{"message": util.MicrosToMillis(final.Duration()), "status": "Success"}
		send("complete", &WorkflowMessage{Idx: -1, Ctx: msg})
	}()
	return "", nil
}

func wireSocketFuncs(
	rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState,
) (func(cmd string, x any), func(i int, s string), func(i int, err error), func(i int, w *loadtoad.WorkflowResult), error) {
	channel := string(rc.URI().QueryArgs().Peek("channel"))
	if channel == "" {
		return nil, nil, nil, nil, errors.New("must provide channel")
	}
	send := func(cmd string, x any) {
		msg := &websocket.Message{Channel: channel, Cmd: cmd, Param: util.ToJSONBytes(x, true)}
		_ = as.Services.Socket.WriteChannel(msg, ps.Logger)
	}
	err := as.Services.Socket.Upgrade(ps.Context, rc, channel, ps.Profile, ps.Logger) //nolint:contextcheck
	if err != nil {
		ps.Logger.Warnf("unable to upgrade connection to WebSocket: %s", err.Error())
		return nil, nil, nil, nil, err
	}
	logF := func(i int, s string) {
		ps.Logger.Infof("[%d] [LOG] %s", i, s)
		send("log", &WorkflowMessage{Idx: i, Ctx: s})
	}
	errF := func(i int, err error) {
		ps.Logger.Warnf("[%d] [ERROR] %+v", i, err)
		send("error", &WorkflowMessage{Idx: i, Ctx: err.Error()})
	}
	okF := func(i int, w *loadtoad.WorkflowResult) {
		ps.Logger.Infof("[%d] [OK] %s", i, w.ID)
		res := vworkflow.RenderResultModal(fmt.Sprintf("result-%d", i), w, ps)
		c := util.ValueMap{"table": vworkflow.RenderResultTable(i, w, ps), "result": res}
		send("ok", &WorkflowMessage{Idx: i, Ctx: c})
	}
	return send, logF, errF, okF, nil
}
