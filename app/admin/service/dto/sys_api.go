package dto

import (
	"go-admin/core/dto"
)

// SysApiQueryReq 功能列表请求参数
type SysApiQueryReq struct {
	dto.Pagination `search:"-"`
	Description    string   `form:"description"  search:"type:contains;column:description;table:sys_api" comment:"功能描述"`
	Path           string   `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Method         string   `form:"method"  search:"type:exact;column:method;table:sys_api" comment:"请求方法"`
	ApiType        string   `form:"apiType"  search:"type:exact;column:api_type;table:sys_api" comment:"接口类型"`
	ApiTypes       []string `form:"apiTypes"  search:"type:in;column:api_type;table:sys_api" comment:"接口类型"`
	BeginCreatedAt string   `form:"beginCreatedAt" search:"type:gte;column:created_at;table:sys_api" comment:"创建时间"`
	EndCreatedAt   string   `form:"endCreatedAt" search:"type:lte;column:created_at;table:sys_api" comment:"创建时间"`
	SysApiOrder
}

type SysApiOrder struct {
	DescriptionOrder string `search:"type:order;column:description;table:sys_api" form:"descriptionOrder"`
	PathOrder        string `search:"type:order;column:path;table:sys_api" form:"pathOrder"`
	CreatedAtOrder   string `search:"type:order;column:created_at;table:sys_api" form:"createdAtOrder"`
}

func (m *SysApiQueryReq) GetNeedSearch() interface{} {
	return *m
}

// SysApiInsertReq 功能创建请求参数
type SysApiInsertReq struct {
	Id          int    `json:"-" comment:"编码"` // 编码
	Description string `json:"description" comment:"功能描述"`
	//Path  string `json:"path" comment:"地址"`
	ApiType    string `json:"apiType" comment:""`
	Method     string `json:"method" comment:"请求方法"`
	CurrUserId int64  `json:"-" comment:""`
}

// SysApiUpdateReq 功能更新请求参数
type SysApiUpdateReq struct {
	Id          int64  `uri:"id" json:"-" comment:"编码"` // 编码
	Description string `json:"description" comment:"功能描述"`
	//Path       string `json:"path" comment:"地址"`
	ApiType string `json:"apiType" comment:""`
	//Method     string `json:"method" comment:"请求方法"`
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
