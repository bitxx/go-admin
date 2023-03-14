package dto

import (
	"go-admin/common/dto"
)

// SysApiQueryReq 功能列表请求参数
type SysApiQueryReq struct {
	dto.Pagination `search:"-"`
	Title          string   `form:"title"  search:"type:contains;column:title;table:sys_api" comment:"标题"`
	Path           string   `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Action         string   `form:"action"  search:"type:exact;column:action;table:sys_api" comment:"请求方法"`
	ApiType        string   `form:"apiType"  search:"type:exact;column:api_type;table:sys_api" comment:"接口类型"`
	ApiTypes       []string `form:"apiTypes"  search:"type:in;column:api_type;table:sys_api" comment:"接口类型"`
	SysApiOrder
}

type SysApiOrder struct {
	TitleOrder     string `search:"type:order;column:title;table:sys_api" form:"titleOrder"`
	PathOrder      string `search:"type:order;column:path;table:sys_api" form:"pathOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_api" form:"createdAtOrder"`
}

func (m *SysApiQueryReq) GetNeedSearch() interface{} {
	return *m
}

// SysApiInsertReq 功能创建请求参数
type SysApiInsertReq struct {
	Id int `json:"-" comment:"编码"` // 编码
	//Handle     string `json:"handle" comment:"handle"`
	Title string `json:"title" comment:"标题"`
	//Path  string `json:"path" comment:"地址"`
	ApiType    string `json:"apiType" comment:""`
	Action     string `json:"action" comment:"类型"`
	CurrUserId int64  `json:"-" comment:""`
}

// SysApiUpdateReq 功能更新请求参数
type SysApiUpdateReq struct {
	Id int64 `uri:"id" json:"-" comment:"编码"` // 编码
	//Handle     string `json:"handle" comment:"handle"`
	Title string `json:"title" comment:"标题"`
	//Path       string `json:"path" comment:"地址"`
	ApiType string `json:"apiType" comment:""`
	//Action     string `json:"action" comment:"类型"`
	CurrUserId int64 `json:"-" comment:"更新者管理员"`
}

// SysApiGetReq 功能获取请求参数
type SysApiGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// SysApiDeleteReq 功能删除请求参数
type SysApiDeleteReq struct {
	Ids []int64 `json:"ids"`
}
