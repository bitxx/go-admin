package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(middleware.Auth())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/pwd/reset", api.ResetPwd)
		r.PUT("/status", api.UpdateStatus)
		r.GET("/logout", api.LogOut)

		r.POST("/profile/avatar", api.InsetProfileAvatar)
		r.PUT("/profile", api.UpdateProfile)
		r.PUT("/profile/pwd", api.UpdateProfilePwd)
		r.GET("/profile", api.GetProfile)
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
