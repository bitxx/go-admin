package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup) {
	api := apis.SysDept{}

	r := v1.Group("/admin/sys/sys-dept").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetTreeList)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/dept-tree", api.GetTree)
		r.GET("/role-dept-tree-select/:roleId", api.GetDeptTreeByRole)
	}

}
