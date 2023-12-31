package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views"
)

var homeContent = util.ValueMap{
	util.AppKey: util.AppName,
	"urls": map[string]string{
		"home": "/",
	},
}

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		hars := as.Services.Har.List(ps.Logger)
		w, _ := as.Services.LoadToad.ListWorkflows(ps.Logger)
		s := as.Services.Script.ListScripts(ps.Logger)
		ps.Data = homeContent
		return Render(rc, as, &views.Home{Hars: hars, Workflows: w, Scripts: s}, ps)
	})
}
