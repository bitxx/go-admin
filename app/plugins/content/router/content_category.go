package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/content/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerContentCategoryRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckContentCategoryRouter)
}

// registerContentCategoryRouter 注册需要认证的路由
func registerContentCategoryRouter(v1 *gin.RouterGroup) {
	api := apis.ContentCategory{}
	r := v1.Group("/plugins/content/content-category").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckContentCategoryRouter
func registerNoCheckContentCategoryRouter(v1 *gin.RouterGroup) {}
