package models

import (
	"time"
)

type UserOperLog struct {
	ActionType string     `json:"actionType" gorm:"column:action_type;type:char(2);comment:用户行为类型"`
	ByType     string     `json:"byType" gorm:"column:by_type;type:char(2);comment:更新用户类型 1-app用户 2-后台用户"`
	CreateBy   int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	Id         int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Remark     string     `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	Status     string     `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-异常)"`
	UpdateBy   int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
	UserId     int64      `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`

	User *User `json:"user" gorm:"foreignkey:user_id"`
}

func (UserOperLog) TableName() string {
	return "app_user_oper_log"
}
