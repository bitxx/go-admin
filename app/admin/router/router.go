package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
	"go-admin/common/runtime"
	"go-admin/common/ws"
	"go-admin/config/config"
	"mime"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter
func InitRouter() {
	var r *gin.Engine
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		panic("not found engine...")
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		panic("not found engine...")
	}
	InitSysRouter(r)

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

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	// 需要认证
	sysCheckRoleRouterInit(g)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()
	r.GET("/info", Ping)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	//静态路由
	r.Static(config.ApplicationConfig.FileRootPath, config.ApplicationConfig.FileRootPath)
	r.Static("/static", "./static")
}

func sysCheckRoleRouterInit(r *gin.RouterGroup) {
	wss := r.Group("").Use(middleware.Auth())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}

	v1 := r.Group("/admin-api/v1")
	registerBaseRouter(v1)
}

func registerBaseRouter(v1 *gin.RouterGroup) {
	api := apis.SysMenu{}
	deptApi := apis.SysDept{}
	v1auth := v1.Group("sys").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		//v1.GET("/menuTreeselect", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeselect/:roleId", deptApi.GetDeptTreeRoleSelect)
	}
}
