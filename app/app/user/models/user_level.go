package models

import (
    "time"
)

type UserLevel struct {
    Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
    Name string `json:"name" gorm:"column:name;type:varchar(255);comment:等级名称"`
    LevelType string `json:"levelType" gorm:"column:level_type;type:varchar(10);comment:等级类型"`
    Level int64 `json:"level" gorm:"column:level;type:int;comment:等级"`
    Status string `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-异常)"`
    Remark string `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
    CreateBy int64 `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
    UpdateBy int64 `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
    CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
    UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (UserLevel) TableName() string {
    return "app_user_level"
}
