package models

import "time"

type SysPost struct {
	Id        int64      `gorm:"primaryKey;autoIncrement" json:"id"` //岗位编号
	PostName  string     `gorm:"size:128;" json:"postName"`          //岗位名称
	PostCode  string     `gorm:"size:128;" json:"postCode"`          //岗位代码
	Sort      int        `gorm:"size:4;" json:"sort"`                //岗位排序
	Status    string     `gorm:"size:1;" json:"status"`              //状态
	Remark    string     `gorm:"size:255;" json:"remark"`            //描述
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`

	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
}

func (SysPost) TableName() string {
	return "sys_post"
}
