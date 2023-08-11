package controller

import (
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
)

type WorkflowMessage struct {
	Idx int `json:"idx"`
	Ctx any `json:"ctx,omitempty"`
}

func WorkflowConnect(rc *fasthttp.RequestCtx) {
	Act("workflow.connect", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := loadWorkflow(as, rc)
		if err != nil {
			return "", err
		}
		channel := string(rc.URI().QueryArgs().Peek("channel"))
		if channel == "" {
			return "", errors.New("must provide channel")
		}
		send := func(cmd string, x any) {
			msg := &websocket.Message{Channel: channel, Cmd: cmd, Param: util.ToJSONBytes(x, true)}
			_ = as.Services.Socket.WriteChannel(msg, ps.Logger)
		}
		err = as.Services.Socket.Upgrade(ps.Context, rc, channel, ps.Profile, ps.Logger)
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
			send("ok", &WorkflowMessage{Idx: i, Ctx: vworkflow.RenderResultTable(i, w, ps)})
		}
		go func() {
			final, e := as.Services.LoadToad.Run(w, logF, errF, okF)
			if e != nil {
				errF(-1, e)
			}
			ps.Logger.Infof("[COMPLETE] %s", w.ID)
			send("complete", &WorkflowMessage{Idx: -1, Ctx: util.MicrosToMillis(final.Duration())})
		}()
		return "", nil
	})
}
