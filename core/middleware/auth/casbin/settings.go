package casbin

import "go-admin/core/global"

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	{Url: global.RouteRootPath + "/v1/dict/type-option-select", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/dict-data/option-select", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/deptTree", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/db/tables/page", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/db/columns/page", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/gen/toproject/:tableId", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/gen/todb/:tableId", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/gen/tabletree", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/gen/preview/:tableId", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/getCaptcha", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/menuTreeselect", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/menurole", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/menuids", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/roleMenuTreeselect/:roleId", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/roleDeptTreeselect/:roleId", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/refresh_token", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/configKey/:configKey", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/app-config", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/login", Method: "POST"},
	{Url: global.RouteRootPath + "/v1/metrics", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/health", Method: "GET"},
	{Url: "/", Method: "GET"},
	{Url: global.RouteRootPath + "/v1/server-monitor", Method: "GET"},
}
