package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/clib"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
)

func makeRoute(x *mux.Router, method string, path string, f http.HandlerFunc) {
	cutil.AddRoute(method, path)
	x.HandleFunc(path, f).Methods(method)
}

func AppRoutes(as *app.State, logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, "/", controller.Home)
	makeRoute(r, http.MethodGet, "/healthcheck", clib.Healthcheck)
	makeRoute(r, http.MethodGet, "/about", clib.About)

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.Profile)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, cutil.DefaultSearchPath, clib.Search)

	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	harRoutes(r)
	loadtoadRoutes(r)
	// $PF_SECTION_END(routes)$

	execRoutes(r)
	scriptingRoutes(r)
	adminRoutes(r)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodOptions, "/", controller.Options)

	return cutil.WireRouter(r, controller.NotFoundAction, logger)
}
