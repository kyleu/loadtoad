// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/loadtoad/app/controller/clib"
)

func harRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/har", clib.HarList)
	makeRoute(r, http.MethodPost, "/har", clib.HarUpload)
	makeRoute(r, http.MethodGet, "/har/{key}", clib.HarDetail)
	makeRoute(r, http.MethodGet, "/har/{key}/delete", clib.HarDelete)
	makeRoute(r, http.MethodGet, "/har/{key}/trim", clib.HarTrim)
}
