package models

import (
    "time"
)

type UserCountryCode struct {
        Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
        Country string `json:"country" gorm:"column:country;type:varchar(64);comment:国家地区"`
        Code string `json:"code" gorm:"column:code;type:varchar(12);comment:区号"`
        Status string `json:"status" gorm:"column:status;type:char(1);comment:状态(1-可用 2-停用)"`
        Remark string `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
        CreateBy int64 `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
        UpdateBy int64 `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
        CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
        UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (UserCountryCode) TableName() string {
    return "app_user_country_code"
}
