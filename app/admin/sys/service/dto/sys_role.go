package dto

import (
	"go-admin/app/admin/sys/models"
	"go-admin/core/dto"
)

// SysRoleQueryReq 列表或者搜索使用结构体
type SysRoleQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id" search:"type:exact;column:id;table:admin_sys_role" comment:"角色编码"`                 // 角色编码
	RoleName       string `form:"roleName" search:"type:contains;column:role_name;table:admin_sys_role" comment:"角色名称"` // 角色名称
	Status         string `form:"status" search:"type:exact;column:status;table:admin_sys_role" comment:"状态 1-正常 2-停用"` // 状态
	RoleKey        string `form:"roleKey" search:"type:exact;column:role_key;table:admin_sys_role" comment:"角色代码"`      // 角色代码
}

type SysRoleOrder struct {
	IdOrder        string `search:"type:order;column:id;table:admin_sys_role" form:"idOrder"`
	RoleNameOrder  string `search:"type:order;column:role_name;table:admin_sys_role" form:"roleNameOrder"`
	RoleSortOrder  string `search:"type:order;column:role_sort;table:admin_sys_role" form:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:admin_sys_role" form:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:admin_sys_role" form:"createdAtOrder"`
}

func (m *SysRoleQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysRoleInsertReq struct {
	RoleName   string           `form:"roleName" comment:"角色名称"`       // 角色名称
	Status     string           `form:"status" comment:"状态 1-正常 2-停用"` // 状态
	RoleKey    string           `form:"roleKey" comment:"角色代码"`        // 角色代码
	RoleSort   int              `form:"roleSort" comment:"角色排序"`       // 角色排序
	Remark     string           `form:"remark" comment:"备注"`           // 备注
	DataScope  string           `form:"dataScope"`
	SysMenu    []models.SysMenu `form:"sysMenu"`
	MenuIds    []int64          `form:"menuIds"`
	SysDept    []models.SysDept `form:"sysDept"`
	DeptIds    []int            `form:"deptIds"`
	CurrUserId int64            `json:"-" comment:""`
}

type SysRoleUpdateReq struct {
	Id         int64   `uri:"id" json:"-" comment:"角色编码"` // 角色编码
	RoleName   string  `form:"roleName" comment:"角色名称"`   // 角色名称
	Status     string  `form:"status" comment:"状态"`       // 状态
	RoleKey    string  `form:"roleKey" comment:"角色代码"`    // 角色代码
	RoleSort   int     `form:"roleSort" comment:"角色排序"`   // 角色排序
	Remark     string  `form:"remark" comment:"备注"`       // 备注
	MenuIds    []int64 `form:"menuIds"`
	CurrUserId int64   `json:"-" comment:""`
}

type UpdateStatusReq struct {
	RoleId     int64  `form:"roleId" comment:"角色编码"` // 角色编码
	Status     string `form:"status" comment:"状态"`   // 状态
	CurrUserId int64  `json:"-" comment:""`
}

// SysRoleGetReq 获取单个或者删除的结构体
type SysRoleGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// RoleDataScopeReq 角色数据权限修改
type RoleDataScopeReq struct {
	Id         int64   `json:"id"`
	DataScope  string  `json:"dataScope"`
	DeptIds    []int64 `json:"deptIds"`
	CurrUserId int64   `json:"-" comment:""`
}

// SysRoleDeleteReq 功能删除请求参数
type SysRoleDeleteReq struct {
	Ids []int64 `json:"ids"`
}
