package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/core"
	"go-admin/common/middleware/auth/jwtauth"
)

const (
	JwtTokenCheck   string = "JwtToken"
	RoleCheck       string = "AuthCheckRole"
	PermissionCheck string = "PermissionAction"
)

func InitMiddleware(r *gin.Engine) {
	// 数据库链接
	r.Use(WithContextDb)
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// KeepAlive is a middleware function that appends headers
	r.Use(KeepAlive)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	//r.Use(DemoEvn())
	// 链路追踪
	//r.Use(middleware.Trace())
	core.Runtime.SetMiddleware(JwtTokenCheck, (*jwtauth.GinJWTMiddleware).MiddlewareFunc)
	core.Runtime.SetMiddleware(RoleCheck, AuthCheckRole())
	core.Runtime.SetMiddleware(PermissionCheck, PermissionAction())
}
