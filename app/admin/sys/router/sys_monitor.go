package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-admin/app/admin/sys/apis"
	"go-admin/core/middleware"
	"go-admin/core/utils/ginutils"
	"net/http"
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
		r.GET("/prom", ginutils.Handler(promhttp.Handler()))
		//健康检查
		r.GET("/health", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}

}
