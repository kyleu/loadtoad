// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/clib"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/loadtoad/app/util"
)

func SiteRoutes(logger util.Logger) fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Site)

	r.GET(cutil.DefaultProfilePath, clib.ProfileSite)
	r.POST(cutil.DefaultProfilePath, clib.ProfileSave)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/assets/{_:*}", clib.Static)

	r.GET("/{path:*}", controller.Site)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFound

	p := httpmetrics.NewMetrics("marketing_site", logger)
	return fasthttp.CompressHandlerLevel(p.WrapHandler(r, false), fasthttp.CompressBestSpeed)
}
