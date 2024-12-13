package models

import "time"

type SysDictData struct {
	Id        int64      `json:"id" gorm:"primaryKey;column:id;autoIncrement;comment:主键编码"`
	DictSort  int        `json:"dictSort" gorm:"size:20;comment:DictSort"`
	DictLabel string     `json:"dictLabel" gorm:"size:128;comment:DictLabel"`
	DictValue string     `json:"dictValue" gorm:"size:255;comment:DictValue"`
	DictType  string     `json:"dictType" gorm:"size:64;comment:DictType"`
	CssClass  string     `json:"cssClass" gorm:"size:128;comment:CssClass"`
	ListClass string     `json:"listClass" gorm:"size:128;comment:ListClass"`
	IsDefault string     `json:"isDefault" gorm:"size:8;comment:IsDefault"`
	Status    string     `json:"status" gorm:"size:4;comment:status"`
	Default   string     `json:"default" gorm:"size:8;comment:Default"`
	Remark    string     `json:"remark" gorm:"size:255;comment:Remark"`
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysDictData) TableName() string {
	return "admin_sys_dict_data"
}
