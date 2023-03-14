package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRuntimeConfigRouter)
}

// 需认证的路由代码
func registerSysRuntimeConfigRouter(v1 *gin.RouterGroup) {
	api := apis.SysRuntimeConfig{}
	r := v1.Group("/sysRuntimeConfig").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("/getConfig", api.GetConfig)
	}
}
