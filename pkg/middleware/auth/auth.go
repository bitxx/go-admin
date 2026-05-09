package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/pkg/middleware/auth/jwtauth"
)

var Auth AuthInter

type AuthInter interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
	GetUserId(c *gin.Context) (int64, int, error)
	GetRoleKey(c *gin.Context) string
	AuthMiddlewareFunc() gin.HandlerFunc
	AuthCheckRoleMiddlewareFunc() gin.HandlerFunc
}

// InitAuth
// @Description: 初始化
func InitAuth() {
	auth, err := jwtauth.NewJwtAuth()
	if err != nil {
		panic(fmt.Sprintf("auth Init Error, %s", err.Error()))
	}
	Auth = auth
}
