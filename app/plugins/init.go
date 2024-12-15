package app

/**
 * app应用体量过大的时候，可以在该应用根目录下新建这样一个文件，将子模块的路由和多语言都辉总在这里，然后再去项目根目录(go-admin/app/init.go)中注册整个应用
 */
import (
	contentRouter "go-admin/app/plugins/content/router"
	filemgrRouter "go-admin/app/plugins/filemgr/router"
	msgRouter "go-admin/app/plugins/msg/router"
)

// AllRouter
// @Description: 汇总各大板块接口
// @return []func()
func AllRouter() []func() {
	//初始化路由
	var routers []func()

	//plugins-消息管理
	routers = append(routers, msgRouter.InitRouter)
	//plugins-内容板块
	routers = append(routers, contentRouter.InitRouter)
	//plugins-文件管理
	routers = append(routers, filemgrRouter.InitRouter)
	return routers
}
