// Package app
// @Description:路由汇总，如无特殊必要，勿操作本页代码
package app

import (
	adminLang "go-admin/app/admin/lang"
	adminRouter "go-admin/app/admin/router"
	appLang "go-admin/app/app"
	appRouter "go-admin/app/app"
	contentLang "go-admin/app/plugins/content/lang"
	contentRouter "go-admin/app/plugins/content/router"
	fileMgrLang "go-admin/app/plugins/filemgr/lang"
	filemgrRouter "go-admin/app/plugins/filemgr/router"
	msgLang "go-admin/app/plugins/msg/lang"
	msgRouter "go-admin/app/plugins/msg/router"
)

// AllRouter
// @Description: 汇总各大板块接口
// @return []func()
func AllRouter() []func() {
	//初始化路由
	var routers []func()

	//app-应用
	routers = append(routers, appRouter.AllRouter()...)

	//plugins-消息管理
	routers = append(routers, msgRouter.InitRouter)
	//plugins-内容板块
	routers = append(routers, contentRouter.InitRouter)
	//plugins-文件管理
	routers = append(routers, filemgrRouter.InitRouter)
	//admin-后台基础管理服务
	routers = append(routers, adminRouter.InitRouter)
	return routers
}

// AllLang
// @Description: 多语言初始化
func AllLang() {
	//初始化多语言
	//app应用
	appLang.LangInit()
	//plugins-消息管理
	msgLang.Init()
	//plugins-文件存储
	fileMgrLang.Init()
	//plugins-内容板块
	contentLang.Init()
	//后台基础
	adminLang.Init()
}
