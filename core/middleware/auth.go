package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/core/middleware/auth"
)

func Auth() gin.HandlerFunc {
	return auth.Auth.AuthMiddlewareFunc()
}

func AuthCheckRole() gin.HandlerFunc {
	return auth.Auth.AuthCheckRoleMiddlewareFunc()
}
