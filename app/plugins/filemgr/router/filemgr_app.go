package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/filemgr/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFilemgrAppRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckFilemgrAppRouter)
}

// registerFilemgrAppRouter 注册需要认证的路由
func registerFilemgrAppRouter(v1 *gin.RouterGroup) {
	api := apis.FilemgrApp{}
	r := v1.Group("/plugins/filemgr/filemgr-app").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.POST("/upload", api.Upload)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckFilemgrAppRouter
func registerNoCheckFilemgrAppRouter(v1 *gin.RouterGroup) {}
