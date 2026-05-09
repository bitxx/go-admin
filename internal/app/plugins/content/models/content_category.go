package models

import (
	"time"
)

type ContentCategory struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name      string     `json:"name" gorm:"column:name;type:varchar(255);comment:分类名称"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:状态（1-正常 2-异常）"`
	Remark    string     `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:更新人编号"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ContentCategory) TableName() string {
	return "plugins_content_category"
}
