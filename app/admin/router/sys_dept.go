package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup) {
	api := apis.SysDept{}

	r := v1.Group("/sys/dept").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetList)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("sys").Use(middleware.Auth())
	{
		r1.GET("/deptTree", api.Get2Tree)
	}

}
