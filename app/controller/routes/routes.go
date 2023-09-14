// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/clib"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/loadtoad/app/util"
)

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Home)
	r.GET("/healthcheck", clib.Healthcheck)
	r.GET("/about", clib.About)

	r.GET(cutil.DefaultProfilePath, clib.Profile)
	r.POST(cutil.DefaultProfilePath, clib.ProfileSave)
	r.GET(cutil.DefaultSearchPath, clib.Search)
	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	harRoutes(r)
	loadtoadRoutes(r)
	// $PF_SECTION_END(routes)$

	r.GET("/admin", clib.Admin)
	execRoutes(r)
	scriptingRoutes(r)
	r.GET("/admin/{path:*}", clib.Admin)
	r.POST("/admin/{path:*}", clib.Admin)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/robots.txt", clib.RobotsTxt)
	r.GET("/assets/{_:*}", clib.Static)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFound

	clib.AppRoutesList = r.List()

	p := httpmetrics.NewMetrics(util.AppKey, logger)
	return fasthttp.CompressHandlerLevel(p.WrapHandler(r, true), fasthttp.CompressBestSpeed)
}
