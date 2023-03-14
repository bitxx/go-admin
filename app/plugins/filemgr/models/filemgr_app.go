package models

import (
    "time"
)

type FilemgrApp struct {
    Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
    Version string `json:"version" gorm:"column:version;type:varchar(100);comment:版本号"`
    Platform string `json:"platform" gorm:"column:platform;type:char(1);comment:平台 (1-安卓 2-苹果)"`
    AppType string `json:"appType" gorm:"column:app_type;type:char(1);comment:版本(1-默认)"`
    LocalAddress string `json:"localAddress" gorm:"column:local_address;type:varchar(255);comment:本地地址"`
    DownloadNum int64 `json:"downloadNum" gorm:"column:download_num;type:int;comment:下载数量"`
    DownloadType string `json:"downloadType" gorm:"column:download_type;type:char(1);comment:下载类型(1-本地 2-外链 3-oss )"`
    DownloadUrl string `json:"downloadUrl" gorm:"column:download_url;type:varchar(255);comment:下载地址(download_type=1使用)"`
    Remark string `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
    Status string `json:"status" gorm:"column:status;type:char(1);comment:状态（1-已发布 2-待发布）
"`
    CreateBy int64 `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
    CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
    UpdateBy int64 `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
    UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (FilemgrApp) TableName() string {
    return "plugins_filemgr_app"
}
