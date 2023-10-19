package router

import (
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"go-admin/common"
	"os"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter 初始化路由
func InitRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		logbase.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		logbase.Fatal("not support other engine")
		os.Exit(-1)
	}

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r)
}

// noCheckRoleRouter 无需认证的路由示例
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/admin-api/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// checkRoleRouter 需要认证的路由示例
func checkRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/admin-api/v1")

	for _, f := range routerCheckRole {
		f(v1)
	}
}
