package controller

import (
	"net/http"

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

func Home(w http.ResponseWriter, r *http.Request) {
	Act("home", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		hars := as.Services.Har.List(ps.Logger)
		wf, _ := as.Services.LoadToad.ListWorkflows(ps.Logger)
		s := as.Services.Script.ListScripts(ps.Logger)
		ps.Data = homeContent
		return Render(r, as, &views.Home{Hars: hars, Workflows: wf, Scripts: s}, ps)
	})
}
