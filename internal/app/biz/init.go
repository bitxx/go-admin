package biz

/**
 * app应用体量过大的时候，可以在该应用根目录下新建这样一个文件，将子模块的路由和多语言都汇总在这里，然后再去项目根目录(go-admin/internal/app/init.go)中注册整个应用
 */
import (
	userRouter "go-admin/internal/app/biz/user/router"
)

// AllRouter
// @Description: 汇总各大板块接口
// @return []func()
func AllRouter() []func() {
	//初始化路由
	var routers []func()

	//biz-用户管理
	routers = append(routers, userRouter.InitRouter)
	return routers
}
