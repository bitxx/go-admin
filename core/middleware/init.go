package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/core/middleware/auth/jwtauth"
	"go-admin/core/runtime"
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
	// IsKeepAlive is a middleware function that appends headers
	r.Use(KeepAlive)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// 链路追踪
	r.Use(Trace())
	runtime.RuntimeConfig.SetMiddleware(JwtTokenCheck, (*jwtauth.GinJWTMiddleware).MiddlewareFunc)
	runtime.RuntimeConfig.SetMiddleware(RoleCheck, AuthCheckRole())
	runtime.RuntimeConfig.SetMiddleware(PermissionCheck, PermissionAction())
}
