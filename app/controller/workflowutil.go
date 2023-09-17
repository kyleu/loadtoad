package controller

import (
	"fmt"
	"slices"
	"strings"

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

var benchArgs = cutil.Args{
	{Key: "concurrency", Title: "Concurrent Runners", Description: "The number of workflow runners that will execute at once", Default: "1"},
	{Key: "count", Title: "Test Count (per runner)", Description: "The number of workflow runs that will execute for each runner", Default: "1"},
}

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

func collectArgs(key string, w *loadtoad.Workflow, rc *fasthttp.RequestCtx) *cutil.ArgResults {
	var args cutil.Args
	if key == "bench" {
		args = slices.Clone(benchArgs)
	}
	for k, v := range w.Replacements {
		if strings.Contains(v, "||") {
			choices := util.StringSplitAndTrim(v, "||")
			args = append(args, &cutil.Arg{Key: k, Title: k, Type: "string", Default: "", Choices: choices})
		} else {
			args = append(args, &cutil.Arg{Key: k, Title: k, Type: "string", Default: v})
		}
	}
	args = append(args, &cutil.Arg{Key: "variables", Title: "Other Variables", Type: "textarea", Default: util.ToJSON(w.Variables)})
	return cutil.CollectArgs(rc, args)
}
