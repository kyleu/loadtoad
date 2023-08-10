package controller

import (
	"github.com/kyleu/loadtoad/views"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
)

func HarList(rc *fasthttp.RequestCtx) {
	Act("har.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := "TODO"
		ps.Data = ret
		return Render(rc, as, &views.Debug{}, ps)
	})
}

func HarDetail(rc *fasthttp.RequestCtx) {
	Act("har.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := "TODO"
		ps.Data = ret
		return Render(rc, as, &views.Debug{}, ps)
	})
}
