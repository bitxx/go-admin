// Package app
// @Description:路由汇总，如无特殊必要，勿操作本页代码
package app

import (
	adminRouter "go-admin/internal/app/admin"
	bizRouter "go-admin/internal/app/biz"
	pluginsRouter "go-admin/internal/app/plugins"
)

// AllRouter
// @Description: 汇总各大板块接口
// @return []func()
func AllRouter() []func() {
	//初始化路由
	var routers []func()

	//biz-应用
	routers = append(routers, bizRouter.AllRouter()...)

	//admin-基础服务
	routers = append(routers, adminRouter.AllRouter()...)

	//plugins-插件服务
	routers = append(routers, pluginsRouter.AllRouter()...)

	return routers
}
