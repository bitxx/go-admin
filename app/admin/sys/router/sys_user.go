package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup) {
	api := apis.SysUser{}
	r := v1.Group("/admin/sys/sys-user").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/pwd-reset", api.ResetPwd)
		r.PUT("/update-status", api.UpdateStatus)
		r.GET("/logout", api.LogOut)
		r.GET("/profile", api.GetProfile)
		r.PUT("/profile", api.UpdateProfile)
		r.POST("/profile/avatar", api.UpdateProfileAvatar)

		r.PUT("/profile/pwd", api.UpdateProfilePwd)
	}
}

func registerNoCheckUserRouter(v1 *gin.RouterGroup) {
	api := apis.SysUser{}
	r := v1.Group("")
	{
		r.GET("/captcha", api.GenCaptcha)
		r.POST("/login", api.Login)
	}
}
