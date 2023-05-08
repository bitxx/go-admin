package models

import "time"

type SysMenu struct {
	Id         int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string     `json:"name" gorm:"size:128;"`
	Title      string     `json:"title" gorm:"size:128;"`
	Icon       string     `json:"icon" gorm:"size:128;"`
	Path       string     `json:"path" gorm:"size:128;"`
	MenuType   string     `json:"menuType" gorm:"size:1;"`
	Permission string     `json:"permission" gorm:"size:255;"`
	ParentIds  string     `json:"parentIds" gorm:"size:255;"`
	ParentId   int64      `json:"parentId" gorm:"size:11;"`
	KeepAlive  string     `json:"keepAlive" gorm:"size:1;"`
	Breadcrumb string     `json:"breadcrumb" gorm:"size:255;"`
	Component  string     `json:"component" gorm:"size:255;"`
	Sort       int        `json:"sort" gorm:"size:4;"`
	Hidden     string     `json:"hidden" gorm:"size:1;comment:1-隐藏 2-显示"`
	IsFrame    string     `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	SysApi     []SysApi   `json:"sysApi" gorm:"many2many:sys_menu_api_rule;foreignKey:id;joinForeignKey:sys_menu_menu_id;references:Id;joinReferences:sys_api_id;comment:外链 1-是 2-否"`
	CreateBy   int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy   int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	Apis       []int      `json:"apis" gorm:"-"`
	DataScope  string     `json:"dataScope" gorm:"-"`
	Params     string     `json:"params" gorm:"-"`
	RoleId     int64      `gorm:"-"`
	Children   []SysMenu  `json:"children,omitempty" gorm:"-"`
	IsSelect   bool       `json:"is_select" gorm:"-"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
