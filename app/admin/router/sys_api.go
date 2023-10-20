package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysApiRouter)
}

// registerSysApiRouter
func registerSysApiRouter(v1 *gin.RouterGroup) {
	api := apis.SysApi{}
	r := v1.Group("/sys-api").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
		r.GET("/export", api.Export)
		r.DELETE("", api.Delete)
	}
}
