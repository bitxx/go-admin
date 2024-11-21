package dto

import (
	"go-admin/app/admin/models"
	"go-admin/core/dto"
)

// SysMenuQueryReq 列表或者搜索使用结构体
type SysMenuQueryReq struct {
	dto.Pagination `search:"-"`
	Title          string  `form:"title" search:"type:contains;column:title;table:sys_menu" comment:"菜单名称"`          // 菜单名称
	Hidden         bool    `form:"hidden" search:"type:exact;column:hidden;table:sys_menu" comment:"显示状态 1-隐藏 2-显示"` // 显示状态
	ParentId       int64   `form:"-" search:"type:exact;column:parent_id;table:sys_menu" comment:"父级"`               // 父级
	ParentIds      []int64 `form:"-" search:"type:in;column:parent_id;table:sys_menu" comment:"父级"`                  // 父级
	MenuIds        []int64 `form:"-" search:"type:in;column:id;table:sys_menu" comment:"菜单编号"`                       // 菜单编号
}

func (m *SysMenuQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysMenuInsertReq struct {
	Name       string          `form:"name" comment:"菜单name"`   //菜单name
	Title      string          `form:"title" comment:"显示名称"`    //显示名称
	Icon       string          `form:"icon" comment:"图标"`       //图标
	Path       string          `form:"path" comment:"路径"`       //路径
	Redirect   string          `form:"redirect" comment:"跳转"`   //针对目录跳转，比如搜索出菜单
	Element    string          `form:"element" comment:"组件"`    //组件
	MenuType   string          `form:"menuType" comment:"菜单类型"` //菜单类型
	SysApi     []models.SysApi `form:"sysApi"`
	Apis       []int           `form:"apis"`
	Permission string          `form:"permission" comment:"权限编码"`        //权限编码
	ParentId   int64           `form:"parentId" comment:"上级菜单"`          //上级菜单
	KeepAlive  string          `form:"keepAlive" comment:"是否缓存 1-是 2-否"` //是否缓存
	IsAffix    string          `form:"isAffix" comment:"是否固定 1-是 2-否"`   //是否固定
	Sort       int             `form:"sort" comment:"排序"`                //排序
	Hidden     string          `form:"hidden" comment:"1-隐藏 2-显示"`       //是否显示
	IsFrame    string          `form:"isFrame" comment:"外链 1-是 2-否"`     //是否frame
	CurrUserId int64           `json:"-" comment:""`
}

type SysMenuUpdateReq struct {
	Id         int64           `uri:"id" json:"-" comment:"编码"` // 编码
	Name       string          `form:"name" comment:"菜单name"`   //菜单name
	Title      string          `form:"title" comment:"显示名称"`    //显示名称
	Icon       string          `form:"icon" comment:"图标"`       //图标
	Path       string          `form:"path" comment:"路径"`       //路径
	Redirect   string          `form:"redirect" comment:"跳转"`   //针对目录跳转，比如搜索出菜单
	Element    string          `form:"element" comment:"组件"`    //组件
	MenuType   string          `form:"menuType" comment:"菜单类型"` //菜单类型
	SysApi     []models.SysApi `form:"sysApi"`
	Apis       []int           `form:"apis"`
	Permission string          `form:"permission" comment:"权限编码"`        //权限编码
	ParentId   int64           `form:"parentId" comment:"上级菜单"`          //上级菜单
	KeepAlive  string          `form:"keepAlive" comment:"是否缓存 1-是 2-否"` //是否缓存
	IsAffix    string          `form:"isAffix" comment:"是否固定 1-是 2-否"`   //是否固定
	Sort       int             `form:"sort" comment:"排序"`                //排序
	Hidden     string          `form:"hidden" comment:"1-隐藏 2-显示"`       //是否显示
	IsFrame    string          `form:"isFrame" comment:"外链 1-是 2-否"`     //是否frame
	CurrUserId int64           `json:"-" comment:""`
}

type SysMenuDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type SysMenuGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

type MenuLabel struct {
	Id       int64       `json:"id,omitempty" gorm:"-"`
	Label    string      `json:"label,omitempty" gorm:"-"`
	Children []MenuLabel `json:"children,omitempty" gorm:"-"`
}

type SelectMenuRole struct {
	RoleId int64 `uri:"roleId"`
}

type MenuTreeRoleResp struct {
	Menus       []MenuLabel `json:"menus"`
	CheckedKeys []int64     `json:"checkedKeys"`
}
