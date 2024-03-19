package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/loadtoad/app/controller"
)

func loadtoadRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/har/{key}/run", controller.HarStartRun)
	makeRoute(r, http.MethodGet, "/har/{key}/bench", controller.HarStartBench)
	makeRoute(r, http.MethodGet, "/har/{key}/connect", controller.HarConnect)

	makeRoute(r, http.MethodGet, "/workflow", controller.WorkflowList)
	makeRoute(r, http.MethodGet, "/workflow/new", controller.WorkflowNew)
	makeRoute(r, http.MethodPost, "/workflow/new", controller.WorkflowCreate)
	makeRoute(r, http.MethodGet, "/workflow/{key}", controller.WorkflowDetail)
	makeRoute(r, http.MethodGet, "/workflow/{key}/edit", controller.WorkflowForm)
	makeRoute(r, http.MethodPost, "/workflow/{key}/edit", controller.WorkflowSave)
	makeRoute(r, http.MethodGet, "/workflow/{key}/delete", controller.WorkflowDelete)
	makeRoute(r, http.MethodGet, "/workflow/{key}/run", controller.WorkflowStartRun)
	makeRoute(r, http.MethodGet, "/workflow/{key}/run/connect", controller.WorkflowConnectRun)
	makeRoute(r, http.MethodGet, "/workflow/{key}/bench", controller.WorkflowStartBench)
	makeRoute(r, http.MethodGet, "/workflow/{key}/bench/connect", controller.WorkflowConnectBench)
}
