package dto

import (
	"github.com/shopspring/decimal"
	"go-admin/common/dto"
)

// SysDeptQueryReq 列表或者搜索使用结构体
type SysDeptQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id" search:"type:exact;column:dept_id;table:sys_dept" comment:"id"`                  //id
	ParentId       int    `form:"parentId" search:"type:exact;column:parent_id;table:sys_dept" comment:"上级部门"`    //上级部门
	DeptPath       string `form:"-" search:"type:contains;column:dept_path;table:sys_dept" comment:""`                //路径
	DeptName       string `form:"deptName" search:"type:contains;column:dept_name;table:sys_dept" comment:"部门名称"` //部门名称
	Sort           int    `form:"-" search:"type:exact;column:sort;table:sys_dept" comment:"排序"`                    //排序
	Leader         string `form:"leader" search:"type:contains;column:leader;table:sys_dept" comment:"负责人"`        //负责人
	Phone          string `form:"phone" search:"type:exact;column:phone;table:sys_dept" comment:"手机"`               //手机
	Email          string `form:"email" search:"type:exact;column:email;table:sys_dept" comment:"邮箱"`               //邮箱
	Status         string `form:"-" search:"type:exact;column:status;table:sys_dept" comment:"状态"`                  //状态
}

func (m *SysDeptQueryReq) GetNeedSearch() interface{} {
	return *m
}

// SysDeptInsertReq 增、改使用的结构体
type SysDeptInsertReq struct {
	ParentId   decimal.Decimal `json:"parentId" comment:"上级部门"` //上级部门
	DeptPath   string          `json:"deptPath" comment:""`         //路径
	DeptName   string          `json:"deptName" comment:"部门名称"` //部门名称
	Sort       int             `json:"sort" comment:"排序"`         //排序
	Leader     string          `json:"leader" comment:"负责人"`     //负责人
	Phone      string          `json:"phone" comment:"手机"`        //手机
	Email      string          `json:"email" comment:"邮箱"`        //邮箱
	CurrUserId int64           `json:"-" comment:""`
}

type SysDeptUpdateReq struct {
	Id         int64  `uri:"id" json:"-" comment:"编码"`   // 编码
	DeptName   string `json:"deptName" comment:"部门名称"` //部门名称
	Sort       int    `json:"sort" comment:"排序"`         //排序
	Leader     string `json:"leader" comment:"负责人"`     //负责人
	Phone      string `json:"phone" comment:"手机"`        //手机
	Email      string `json:"email" comment:"邮箱"`        //邮箱
	CurrUserId int64  `json:"-" comment:""`
}

// SysDeptGetReq 获取单个
type SysDeptGetReq struct {
	Id int64 `uri:"id" json:"-" gorm:""`
}

type SysRoleDeptResp struct {
	DeptId int64 `json:"-"`
}

type SysDeptDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type DeptLabel struct {
	Id       int64       `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLabel `gorm:"-" json:"children"`
}

type SelectDeptRole struct {
	RoleId int64 `uri:"roleId"`
}

type DeptTreeRoleResp struct {
	Depts       []DeptLabel `json:"depts"`
	CheckedKeys []int64     `json:"checkedKeys"`
}
