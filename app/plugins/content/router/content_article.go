package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/content/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerContentArticleRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckContentArticleRouter)
}

// registerContentArticleRouter 注册需要认证的路由
func registerContentArticleRouter(v1 *gin.RouterGroup) {
	api := apis.ContentArticle{}
	r := v1.Group("/plugins/content/content-article").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckContentArticleRouter
func registerNoCheckContentArticleRouter(v1 *gin.RouterGroup) {}
