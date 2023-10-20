package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserLevelRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserLevelRouter)
}

// registerUserLevelRouter 注册需要认证的路由
func registerUserLevelRouter(v1 *gin.RouterGroup) {
	api := apis.UserLevel{}
	r := v1.Group("/app/user/user-level").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserLevelRouter
func registerNoCheckUserLevelRouter(v1 *gin.RouterGroup) {}
