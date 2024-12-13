package dto

import (
	"go-admin/core/dto"
)

// SysPostQueryReq 列表或者搜索使用结构体
type SysPostQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id" search:"type:exact;column:id;table:sys_post" comment:"id"`                 // id
	PostName       string `form:"postName" search:"type:contains;column:post_name;table:sys_post" comment:"名称"` // 名称
	PostCode       string `form:"postCode" search:"type:contains;column:post_code;table:sys_post" comment:"编码"` // 编码
	Sort           int    `form:"sort" search:"type:exact;column:sort;table:sys_post" comment:"排序"`             // 排序
	Status         string `form:"status" search:"type:exact;column:status;table:sys_post" comment:"状态"`         // 状态
	Remark         string `form:"remark" search:"type:exact;column:remark;table:sys_post" comment:"备注"`         // 备注
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:sys_post" comment:"创建时间"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:sys_post" comment:"创建时间"`
}

func (m *SysPostQueryReq) GetNeedSearch() interface{} {
	return *m
}

// SysPostInsertReq 增使用的结构体
type SysPostInsertReq struct {
	PostName   string `form:"postName"  comment:"名称"`
	PostCode   string `form:"postCode" comment:"编码"`
	Sort       int    `form:"sort" comment:"排序"`
	Status     string `form:"status"   comment:"状态"`
	Remark     string `form:"remark"   comment:"备注"`
	CurrUserId int64  `json:"-" comment:""`
}

// SysPostUpdateReq 改使用的结构体
type SysPostUpdateReq struct {
	Id         int64  `uri:"id" json:"-"  comment:"id"`
	PostName   string `form:"postName"  comment:"名称"`
	PostCode   string `form:"postCode" comment:"编码"`
	Sort       int    `form:"sort" comment:"排序"`
	Status     string `form:"status"   comment:"状态"`
	Remark     string `form:"remark"   comment:"备注"`
	CurrUserId int64  `form:"-" comment:""`
}

// SysPostGetReq 获取单个的结构体
type SysPostGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

func (s *SysPostGetReq) GetId() interface{} {
	return s.Id
}

// SysPostDeleteReq 删除的结构体
type SysPostDeleteReq struct {
	Ids []int64 `json:"ids"`
}
