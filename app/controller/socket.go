package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

type WorkflowMessage struct {
	Idx int `json:"idx"`
	Ctx any `json:"ctx,omitempty"`
}

func HarConnect(rc *fasthttp.RequestCtx) {
	Act("har.connect", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.LoadHar(key)
		if err != nil {
			return "", err
		}
		w := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		return socketConnect(ps.Context, w, map[string]string{}, rc, as, ps)
	})
}

func WorkflowConnect(rc *fasthttp.RequestCtx) {
	Act("workflow.connect", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}

		repls := w.Replacements
		if string(rc.URI().QueryArgs().Peek("ok")) != util.BoolTrue {
			_, argRes := collectArgs(w, rc)
			if argRes.HasMissing() {
				return "", errors.Errorf("missing replacements [%s]", strings.Join(argRes.Missing, ", "))
			}
			repls = argRes.Values
		}
		return socketConnect(ps.Context, w, repls, rc, as, ps)
	})
}

func socketConnect(
	ctx context.Context, w *loadtoad.Workflow, repls map[string]string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState,
) (string, error) {
	channel := string(rc.URI().QueryArgs().Peek("channel"))
	if channel == "" {
		return "", errors.New("must provide channel")
	}
	send := func(cmd string, x any) {
		msg := &websocket.Message{Channel: channel, Cmd: cmd, Param: util.ToJSONBytes(x, true)}
		_ = as.Services.Socket.WriteChannel(msg, ps.Logger)
	}
	err := as.Services.Socket.Upgrade(ps.Context, rc, channel, ps.Profile, ps.Logger) //nolint:contextcheck
	if err != nil {
		ps.Logger.Warnf("unable to upgrade connection to WebSocket: %s", err.Error())
		return "", err
	}
	logF := func(i int, s string) {
		ps.Logger.Infof("[%d] [LOG] %s", i, s)
		send("log", &WorkflowMessage{Idx: i, Ctx: s})
	}
	errF := func(i int, err error) {
		ps.Logger.Warnf("[%d] [ERROR] %s", i, err.Error())
		send("error", &WorkflowMessage{Idx: i, Ctx: err.Error()})
	}
	okF := func(i int, w *loadtoad.WorkflowResult) {
		ps.Logger.Infof("[%d] [OK] %s", i, w.ID)
		res := vworkflow.RenderResultModal(fmt.Sprintf("result-%d", i), w, ps)
		c := util.ValueMap{"table": vworkflow.RenderResultTable(i, w, ps), "result": res}
		send("ok", &WorkflowMessage{Idx: i, Ctx: c})
	}
	go func() {
		final, e := as.Services.LoadToad.Run(ctx, w, repls, logF, errF, okF)
		if e != nil {
			errF(-1, e)
		}
		ps.Logger.Infof("[COMPLETE] %s", w.ID)
		msg := map[string]any{"message": util.MicrosToMillis(final.Duration()), "status": "Success"}
		send("complete", &WorkflowMessage{Idx: -1, Ctx: msg})
	}()
	return "", nil
}

func collectArgs(w *loadtoad.Workflow, rc *fasthttp.RequestCtx) (cutil.Args, *cutil.ArgResults) {
	var args cutil.Args
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
