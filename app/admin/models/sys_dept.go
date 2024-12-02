package models

import "time"

type SysDept struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId  int64      `json:"parentId" gorm:""`                    //上级部门
	DeptPath  string     `json:"deptPath" gorm:"size:255;"`           //
	DeptName  string     `json:"deptName"  gorm:"size:128;"`          //部门名称
	Sort      int        `json:"sort" gorm:"size:4;"`                 //排序
	Leader    string     `json:"leader" gorm:"size:128;"`             //负责人
	Phone     string     `json:"phone" gorm:"size:11;"`               //手机
	Email     string     `json:"email" gorm:"size:64;"`               //邮箱
	Status    string     `json:"status" gorm:"size:2;"`               //状态
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	DataScope string     `json:"dataScope" gorm:"-"`
	Params    string     `json:"params" gorm:"-"`
	Children  []SysDept  `json:"children" gorm:"-"`
	IsFlag    bool       `json:"-" gorm:"-"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
