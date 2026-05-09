package models

import (
    "time"
)

type MsgCode struct {
    Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
    UserId int64 `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
    Code string `json:"code" gorm:"column:code;type:varchar(12);comment:验证码"`
    CodeType string `json:"codeType" gorm:"column:code_type;type:char(1);comment:验证码类型 1-邮箱；2-短信"`
    Remark string `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注异常"`
    Status string `json:"status" gorm:"column:status;type:char(1);comment:验证码状态 1-发送成功 2-发送失败"`
    CreateBy int64 `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
    UpdateBy int64 `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
    CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
    UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (MsgCode) TableName() string {
    return "plugins_msg_code"
}
