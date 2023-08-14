package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/loadtoad/app/controller"
)

func loadtoadRoutes(r *router.Router) {
	r.GET("/har", controller.HarList)
	r.POST("/har", controller.HarUpload)
	r.GET("/har/{key}", controller.HarDetail)
	r.GET("/har/{key}/run", controller.HarStart)
	r.GET("/har/{key}/connect", controller.HarConnect)

	r.GET("/workflow", controller.WorkflowList)
	r.GET("/workflow/new", controller.WorkflowNew)
	r.POST("/workflow/new", controller.WorkflowCreate)
	r.GET("/workflow/{key}", controller.WorkflowDetail)
	r.GET("/workflow/{key}/edit", controller.WorkflowForm)
	r.POST("/workflow/{key}/edit", controller.WorkflowSave)
	r.GET("/workflow/{key}/run", controller.WorkflowStart)
	r.GET("/workflow/{key}/connect", controller.WorkflowConnect)
}
