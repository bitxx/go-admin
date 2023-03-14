package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// registerUserRouter 注册需要认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	api := apis.User{}
	r := v1.Group("/app/user/user").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserRouter
func registerNoCheckUserRouter(v1 *gin.RouterGroup) {}
