// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views"
)

func About(rc *fasthttp.RequestCtx) {
	controller.Act("about", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("About "+util.AppName, util.AppName+" v"+as.BuildInfo.Version)
		return controller.Render(rc, as, &views.About{}, ps, "about")
	})
}
