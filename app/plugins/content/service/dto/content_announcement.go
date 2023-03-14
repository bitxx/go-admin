package dto

import (
	"go-admin/common/dto"
	"time"
)

type ContentAnnouncementQueryReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:in;column:title;table:plugins_content_announcement" comment:"标题"`
	Status         string `form:"status"  search:"type:exact;column:status;table:plugins_content_announcement" comment:"状态（0正常 1删除 2停用 3冻结）"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:plugins_content_announcement" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:plugins_content_announcement" comment:"创建时间"`
	ContentAnnouncementOrder
}

type ContentAnnouncementOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:plugins_content_announcement"`
	TitleOrder     string     `form:"titleOrder"  search:"type:order;column:title;table:plugins_content_announcement"`
	ContentOrder   string     `form:"contentOrder"  search:"type:order;column:content;table:plugins_content_announcement"`
	NumOrder       int64      `form:"numOrder"  search:"type:order;column:num;table:plugins_content_announcement"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:plugins_content_announcement"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:plugins_content_announcement"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:plugins_content_announcement"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:plugins_content_announcement"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:plugins_content_announcement"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:plugins_content_announcement"`
}

func (m *ContentAnnouncementQueryReq) GetNeedSearch() interface{} {
	return *m
}

type ContentAnnouncementInsertReq struct {
	Title      string `json:"title" comment:"标题"`
	Content    string `json:"content" comment:"内容"`
	Num        int64  `json:"num" comment:"阅读次数"`
	Remark     string `json:"remark" comment:"备注信息"`
	Status     string `json:"status" comment:"状态（0正常 1删除 2停用 3冻结）"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type ContentAnnouncementUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"公告编号"` // 公告编号
	Title      string `json:"title" comment:"标题"`
	Content    string `json:"content" comment:"内容"`
	Num        int64  `json:"num" comment:"阅读次数"`
	Remark     string `json:"remark" comment:"备注信息"`
	Status     string `json:"status" comment:"状态（0正常 1删除 2停用 3冻结）"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// ContentAnnouncementGetReq 功能获取请求参数
type ContentAnnouncementGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// ContentAnnouncementDeleteReq 功能删除请求参数
type ContentAnnouncementDeleteReq struct {
	Ids []int64 `json:"ids"`
}
