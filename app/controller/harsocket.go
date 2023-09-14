package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func HarConnect(rc *fasthttp.RequestCtx) {
	Act("har.connect", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		w := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		return socketConnect(ps.Context, w, map[string]string{}, rc, as, ps)
	})
}

func HarStartRun(rc *fasthttp.RequestCtx) {
	Act("har.start.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		w := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		channel := "run-" + util.RandomString(16)
		return Render(rc, as, &vworkflow.Start{Workflow: w, Entries: ret.Entries, Channel: channel, Path: "/har/" + ret.Key + "/connect"}, ps, "har", ret.Key, "Run")
	})
}

func HarStartBench(rc *fasthttp.RequestCtx) {
	Act("har.start.bench", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		w := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		channel := "run-" + util.RandomString(16)
		return Render(rc, as, &vworkflow.Start{Workflow: w, Entries: ret.Entries, Channel: channel, Path: "/har/" + ret.Key + "/connect"}, ps, "har", ret.Key, "Benchmark")
	})
}
