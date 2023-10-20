package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserOperLogRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserOperLogRouter)
}

// registerUserOperLogRouter 注册需要认证的路由
func registerUserOperLogRouter(v1 *gin.RouterGroup) {
	api := apis.UserOperLog{}
	r := v1.Group("/app/user/user-oper-log").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserOperLogRouter
func registerNoCheckUserOperLogRouter(v1 *gin.RouterGroup) {}
