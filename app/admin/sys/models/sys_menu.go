package models

import (
	"time"
)

type SysMenu struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" gorm:"size:128;"`
	Icon        string     `json:"icon" gorm:"size:128;"`
	Path        string     `json:"path" gorm:"size:255;"`
	Element     string     `json:"element" gorm:"size:255;"`
	Redirect    string     `json:"redirect" gorm:"size:255;"` //针对目录跳转，比如搜索出菜单
	Permission  string     `json:"permission" gorm:"size:255;"`
	ParentIds   string     `json:"parentIds" gorm:"size:255;"`
	ParentId    int64      `json:"parentId" gorm:"size:11;"`
	Sort        int        `json:"sort" gorm:"size:4;"`
	MenuType    string     `json:"menuType" gorm:"size:1;"`
	IsKeepAlive string     `json:"isKeepAlive" gorm:"size:1;comment:1-是 2-否"`
	IsAffix     string     `json:"isAffix" gorm:"size:1;comment:1-是 2-否"`
	IsHidden    string     `json:"isHidden" gorm:"size:1;comment:1-是 2-否"`
	IsFrame     string     `json:"isFrame" gorm:"size:1;comment:1-是 2-否"`
	SysApi      []SysApi   `json:"sysApi" gorm:"many2many:admin_sys_menu_api_rule;foreignKey:id;joinForeignKey:admin_sys_menu_menu_id;references:Id;joinReferences:admin_sys_api_id;"`
	CreateBy    int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	Apis        []int64    `json:"apis" gorm:"-"`
	Children    []*SysMenu `json:"children,omitempty" gorm:"-"`
}

func (SysMenu) TableName() string {
	return "admin_sys_menu"
}
