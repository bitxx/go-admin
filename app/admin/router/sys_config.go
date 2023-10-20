package router

import (
	"go-admin/app/admin/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup) {
	api := apis.SysConfig{}
	r := v1.Group("/config").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}

	r1 := v1.Group("/configKey").Use(middleware.Auth())
	{
		r1.GET("/:configKey", api.GetSysConfigByKey)
	}

	r2 := v1.Group("/app-config")
	{
		r2.GET("", api.GetSysConfigBySysApp)
	}

}
