package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserConfRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserConfRouter)
}

// registerUserConfRouter 注册需要认证的路由
func registerUserConfRouter(v1 *gin.RouterGroup) {
	api := apis.UserConf{}
	r := v1.Group("/app/user/user-conf").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
	}
}

// registerNoCheckUserConfRouter
func registerNoCheckUserConfRouter(v1 *gin.RouterGroup) {}
