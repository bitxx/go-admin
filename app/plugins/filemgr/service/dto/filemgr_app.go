package dto

import (
	"go-admin/common/dto"
	"time"
)

type FilemgrAppQueryReq struct {
	dto.Pagination `search:"-"`
	Version        string `form:"version"  search:"type:exact;column:version;table:plugins_filemgr_app" comment:"版本号"`
	Platform       string `form:"platform"  search:"type:exact;column:platform;table:plugins_filemgr_app" comment:"平台 (1-安卓 2-苹果)"`
	AppType        string `form:"appType"  search:"type:exact;column:app_type;table:plugins_filemgr_app" comment:"版本(1-默认)"`
	DownloadType   string `form:"downloadType"  search:"type:exact;column:download_type;table:plugins_filemgr_app" comment:"下载类型(1-本地 2-外链 3-oss )"`
	Status         string `form:"status"  search:"type:exact;column:status;table:plugins_filemgr_app" comment:"状态（1-已发布 2-待发布）"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:plugins_filemgr_app" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:plugins_filemgr_app" comment:"创建时间"`
	FilemgrAppOrder
}

type FilemgrAppOrder struct {
	IdOrder           int64      `form:"idOrder"  search:"type:order;column:id;table:plugins_filemgr_app"`
	VersionOrder      string     `form:"versionOrder"  search:"type:order;column:version;table:plugins_filemgr_app"`
	PlatformOrder     string     `form:"platformOrder"  search:"type:order;column:platform;table:plugins_filemgr_app"`
	AppTypeOrder      string     `form:"appTypeOrder"  search:"type:order;column:app_type;table:plugins_filemgr_app"`
	LocalAddressOrder string     `form:"localAddressOrder"  search:"type:order;column:local_address;table:plugins_filemgr_app"`
	DownloadNumOrder  int64      `form:"downloadNumOrder"  search:"type:order;column:download_num;table:plugins_filemgr_app"`
	DownloadTypeOrder string     `form:"downloadTypeOrder"  search:"type:order;column:download_type;table:plugins_filemgr_app"`
	DownloadUrlOrder  string     `form:"downloadUrlOrder"  search:"type:order;column:download_url;table:plugins_filemgr_app"`
	RemarkOrder       string     `form:"remarkOrder"  search:"type:order;column:remark;table:plugins_filemgr_app"`
	StatusOrder       string     `form:"statusOrder"  search:"type:order;column:status;table:plugins_filemgr_app"`
	CreateByOrder     int64      `form:"createByOrder"  search:"type:order;column:create_by;table:plugins_filemgr_app"`
	CreatedAtOrder    *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:plugins_filemgr_app"`
	UpdateByOrder     int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:plugins_filemgr_app"`
	UpdatedAtOrder    *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:plugins_filemgr_app"`
}

func (m *FilemgrAppQueryReq) GetNeedSearch() interface{} {
	return *m
}

type FilemgrAppInsertReq struct {
	Version      string `json:"version" comment:"版本号"`
	Platform     string `json:"platform" comment:"平台 (1-安卓 2-苹果)"`
	AppType      string `json:"appType" comment:"版本(1-默认)"`
	LocalAddress string `json:"localAddress" comment:"本地地址"`
	LocalRootUrl string `json:"localRootUrl" comment:"本地Url根地址"`
	DownloadNum  int64  `json:"downloadNum" comment:"下载数量"`
	DownloadType string `json:"downloadType" comment:"下载类型(1-本地 2-外链 3-oss )"`
	DownloadUrl  string `json:"downloadUrl" comment:"下载地址"`
	Remark       string `json:"remark" comment:"备注信息"`
	CurrUserId   int64  `json:"-" comment:"当前登陆用户"`
}

type FilemgrAppUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"App编号"` // App编号
	Status     string `json:"status" comment:"状态（1-已发布 2-待发布）"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// FilemgrAppGetReq 功能获取请求参数
type FilemgrAppGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// FilemgrAppDeleteReq 功能删除请求参数
type FilemgrAppDeleteReq struct {
	Ids []int64 `json:"ids"`
}
