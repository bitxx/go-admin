package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/app/admin/sys/apis"
	"go-admin/pkg/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMonitorRouter)
}

// 需认证的路由代码
func registerSysMonitorRouter(v1 *gin.RouterGroup) {
	api := apis.Monitor{}
	r := v1.Group("/admin/sys/sys-monitor").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetMonitor)
		r.GET("/ping", api.Ping)
	}

}
