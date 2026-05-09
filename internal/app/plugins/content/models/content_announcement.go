package models

import (
	"time"
)

type ContentAnnouncement struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Title     string     `json:"title" gorm:"column:title;type:varchar(255);comment:标题"`
	Content   string     `json:"content" gorm:"column:content;type:text;comment:内容"`
	Num       int64      `json:"num" gorm:"column:num;type:int;comment:阅读次数"`
	Remark    string     `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:状态状态（1-正常 2-异常）"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ContentAnnouncement) TableName() string {
	return "plugins_content_announcement"
}
