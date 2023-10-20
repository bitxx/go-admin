package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/content/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerContentAnnouncementRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckContentAnnouncementRouter)
}

// registerContentAnnouncementRouter 注册需要认证的路由
func registerContentAnnouncementRouter(v1 *gin.RouterGroup) {
	api := apis.ContentAnnouncement{}
	r := v1.Group("/plugins/content/content-announcement").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckContentAnnouncementRouter
func registerNoCheckContentAnnouncementRouter(v1 *gin.RouterGroup) {}
