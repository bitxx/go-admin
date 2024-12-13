package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/sys/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

// 需认证的路由代码
func registerSysRoleRouter(v1 *gin.RouterGroup) {
	api := apis.SysRole{}
	r := v1.Group("/admin/sys/sys-role").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/list", api.GetList)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/role-status", api.UpdateStatus)
		r.PUT("/role-data-scope", api.UpdateDataScope)
	}
}
