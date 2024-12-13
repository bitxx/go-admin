package casbin

import "go-admin/core/global"

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	//主菜单
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-menu/menu-role", Method: "GET"},

	//字典
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-dict/type/option-select", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-dict/data/select", Method: "GET"},

	//登录和退出
	{Url: global.RouteRootPath + "/v1/login", Method: "POST"},
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-user/logout", Method: "GET"},
	{Url: global.RouteRootPath + "/admin-api/v1/captcha", Method: "GET"},

	//登录用户个人信息
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-user/profile", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-user/profile/pwd", Method: "PUT"},
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-user/profile/avatar", Method: "POST"},
	{Url: global.RouteRootPath + "/v1/admin/sys/sys-user/profile", Method: "PUT"},
	//{Url: "/", Method: "GET"},
}
