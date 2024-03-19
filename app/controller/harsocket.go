package controller

import (
	"net/http"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func HarConnect(w http.ResponseWriter, r *http.Request) {
	Act("har.connect", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(r, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		wf := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		return socketConnect(ps.Context, "run", wf, map[string]string{}, w, r, as, ps)
	})
}

func HarStartRun(w http.ResponseWriter, r *http.Request) {
	Act("har.start.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(r, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		wf := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		channel := "run-" + util.RandomString(16)
		pth := "/har/" + ret.Key + "/connect"
		return Render(w, r, as, &vworkflow.Start{Workflow: wf, Entries: ret.Entries, Channel: channel, Path: pth}, ps, "har", ret.Key, "Run")
	})
}

func HarStartBench(w http.ResponseWriter, r *http.Request) {
	Act("har.start.bench", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(r, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		wf := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		channel := "run-" + util.RandomString(16)
		pth := "/har/" + ret.Key + "/connect"
		return Render(w, r, as, &vworkflow.Start{Workflow: wf, Entries: ret.Entries, Channel: channel, Path: pth}, ps, "har", ret.Key, "Benchmark")
	})
}
