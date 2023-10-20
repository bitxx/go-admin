package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/msg/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerMsgCodeRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckMsgCodeRouter)
}

// registerMsgCodeRouter 注册需要认证的路由
func registerMsgCodeRouter(v1 *gin.RouterGroup) {
	api := apis.MsgCode{}
	r := v1.Group("/plugins/msg/msg-code").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
	}
}

// registerNoCheckMsgCodeRouter
func registerNoCheckMsgCodeRouter(v1 *gin.RouterGroup) {}
