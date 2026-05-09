package models

import "time"

type SysDictType struct {
	Id        int64      `json:"id" gorm:"primaryKey;column:id;autoIncrement;comment:主键编码"`
	DictName  string     `json:"dictName" gorm:"size:128;comment:DictName"`
	DictType  string     `json:"dictType" gorm:"size:128;comment:DictType"`
	Status    string     `json:"status" gorm:"size:4;comment:状态"`
	Remark    string     `json:"remark" gorm:"size:255;comment:Remark"`
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysDictType) TableName() string {
	return "admin_sys_dict_type"
}
