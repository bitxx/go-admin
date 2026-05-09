package models

import "time"

type SysConfig struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ConfigName  string     `json:"configName" gorm:"size:128;comment:ConfigName"`   //
	ConfigKey   string     `json:"configKey" gorm:"size:128;comment:ConfigKey"`     //
	ConfigValue string     `json:"configValue" gorm:"size:255;comment:ConfigValue"` //
	ConfigType  string     `json:"configType" gorm:"size:1;comment:ConfigType"`
	IsFrontend  string     `json:"isFrontend" gorm:"size:1;comment:是否前台"` //
	Remark      string     `json:"remark" gorm:"size:128;comment:Remark"` //
	CreateBy    int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysConfig) TableName() string {
	return "admin_sys_config"
}
