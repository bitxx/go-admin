package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserCountryCodeRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserCountryCodeRouter)
}

// registerUserCountryCodeRouter 注册需要认证的路由
func registerUserCountryCodeRouter(v1 *gin.RouterGroup) {
	api := apis.UserCountryCode{}
	r := v1.Group("/app/user/user-country-code").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserCountryCodeRouter
func registerNoCheckUserCountryCodeRouter(v1 *gin.RouterGroup) {}
