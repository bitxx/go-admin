package casbin

import "go-admin/core/global"

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	//主菜单
	{Url: global.RouteRootPath + "/v1/menu/menurole", Method: "GET"},

	//字典
	{Url: global.RouteRootPath + "/v1/dict/type-option-select", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/dict-data/option-select", Method: "GET"},

	//{Url: global.RouteRootPath + "/v1/sys/dept", Method: "GET"},
	//{Url: global.RouteRootPath + "/admin-api/v1/sys/deptTree", Method: "GET"},
	//{Url: global.RouteRootPath + "/admin-api/v1/config/:configKey", Method: "GET"},
	//登录和退出
	{Url: global.RouteRootPath + "/v1/login", Method: "POST"},
	{Url: global.RouteRootPath + "/v1/sys-user/logout", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/getCaptcha", Method: "GET"},

	//登录用户个人信息
	{Url: global.RouteRootPath + "/v1/sys-user/profile", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/sys-user/profile/pwd", Method: "PUT"},
	{Url: global.RouteRootPath + "/v1/sys-user/profile/avatar", Method: "POST"},
	{Url: global.RouteRootPath + "/v1/sys-user/profile", Method: "PUT"},
	{Url: "/", Method: "GET"},
}
