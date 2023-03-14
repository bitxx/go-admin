package models

import (
	"time"
)

type UserConf struct {
	Id       int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId   int64  `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
	CanLogin string `json:"canLogin" gorm:"column:can_login;type:char(1);comment:1-允许登陆；2-不允许登陆"`
	Remark   string `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	Status   string `json:"status" gorm:"column:status;type:char(1);comment:状态（1-正常 2-异常）
"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
	User      *User      `json:"user" gorm:"foreignkey:user_id"`
}

func (UserConf) TableName() string {
	return "app_user_conf"
}
