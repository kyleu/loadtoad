package controller

import (
	"github.com/kyleu/loadtoad/views/vhar"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
)

func HarList(rc *fasthttp.RequestCtx) {
	Act("har.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := as.Services.LoadToad.ListHars(ps.Logger)
		ps.Data = ret
		ps.Title = "Archives"
		return Render(rc, as, &vhar.List{Hars: ret}, ps, "har")
	})
}

func HarDetail(rc *fasthttp.RequestCtx) {
	Act("har.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.LoadHar(key)
		if err != nil {
			return "", err
		}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		return Render(rc, as, &vhar.Detail{Har: ret}, ps, "har", key)
	})
}
