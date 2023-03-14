package models

import (
	"time"
)

type SysRole struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement"` // 角色编码
	RoleName  string     `json:"roleName" gorm:"size:128;"`          // 角色名称
	Status    string     `json:"status" gorm:"size:1;"`              //
	RoleKey   string     `json:"roleKey" gorm:"size:128;"`           //角色代码
	RoleSort  int        `json:"roleSort" gorm:""`                   //角色排序
	Remark    string     `json:"remark" gorm:"size:255;"`            //备注
	DataScope string     `json:"dataScope" gorm:"size:128;"`
	MenuIds   []int64    `json:"menuIds" gorm:"-"`
	DeptIds   []int64    `json:"deptIds" gorm:"-"`
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	SysDept   []SysDept  `json:"sysDept" gorm:"many2many:sys_role_dept;foreignKey:id;joinForeignKey:role_id;references:id;joinReferences:dept_id;"`
	SysMenu   *[]SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:id;joinForeignKey:role_id;references:Id;joinReferences:menu_id;"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
