package routes

import (
	"github.com/fasthttp/router"
	"github.com/kyleu/loadtoad/app/controller"
)

func loadtoadRoutes(r *router.Router) {
	r.GET("/har", controller.HarList)
	r.POST("/har", controller.HarUpload)
	r.GET("/har/{key}", controller.HarDetail)

	r.GET("/workflow", controller.WorkflowList)
	r.GET("/workflow/{key}", controller.WorkflowDetail)
	r.GET("/workflow/{key}/run", controller.WorkflowStart)
	r.GET("/workflow/{key}/connect", controller.WorkflowConnect)
	r.GET("/workflow/{key}/run/sync", controller.WorkflowRunSync)
}
