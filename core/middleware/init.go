package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
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
	runtime.RuntimeConfig.SetMiddleware(JwtTokenCheck, (*jwt.GinJWTMiddleware).MiddlewareFunc)
	runtime.RuntimeConfig.SetMiddleware(RoleCheck, AuthCheckRole())
	runtime.RuntimeConfig.SetMiddleware(PermissionCheck, PermissionAction())
}
