package dto

import (
	"go-admin/common/dto"
	"time"
)

type ContentArticleQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int64   `form:"id"  search:"type:exact;column:id;table:plugins_content_article" comment:"文章编号"`
	CateId         int64   `form:"cateId"  search:"type:exact;column:cate_id;table:plugins_content_article" comment:"分类编号"`
	CateIds        []int64 `form:"-"  search:"type:in;column:cate_id;table:plugins_content_article" comment:"分类编号集合"`
	Name           string  `form:"name"  search:"type:contains;column:name;table:plugins_content_article" comment:"标题"`
	NameInner      string  `form:"-"  search:"type:exact;column:name;table:plugins_content_article" comment:"标题，跟Name区分，数据库精确查询"`
	BeginTime      string  `form:"beginTime" search:"type:gte;column:created_at;table:plugins_content_article" comment:"创建时间"`
	EndTime        string  `form:"endTime" search:"type:lte;column:created_at;table:plugins_content_article" comment:"创建时间"`
	Status         string  `form:"status"  search:"type:exact;column:status;table:plugins_content_article" comment:"状态（1-正常 2-异常）"`
	ContentArticleOrder
}

type ContentArticleOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:plugins_content_article"`
	CateIdOrder    int64      `form:"cateIdOrder"  search:"type:order;column:cate_id;table:plugins_content_article"`
	NameOrder      string     `form:"nameOrder"  search:"type:order;column:name;table:plugins_content_article"`
	ContentOrder   string     `form:"contentOrder"  search:"type:order;column:content;table:plugins_content_article"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:plugins_content_article"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:plugins_content_article"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:plugins_content_article"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:plugins_content_article"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:plugins_content_article"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:plugins_content_article"`
}

func (m *ContentArticleQueryReq) GetNeedSearch() interface{} {
	return *m
}

type ContentArticleInsertReq struct {
	CateId     int64  `json:"cateId" comment:"分类编号"`
	Name       string `json:"name" comment:"标题"`
	Content    string `json:"content" comment:"内容"`
	Remark     string `json:"remark" comment:"备注信息"`
	Status     string `json:"status" comment:"状态（1-正常 2-异常）"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type ContentArticleUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"文章编号"` // 文章编号
	CateId     int64  `json:"cateId" comment:"分类编号"`
	Name       string `json:"name" comment:"标题"`
	Content    string `json:"content" comment:"内容"`
	Remark     string `json:"remark" comment:"备注信息"`
	Status     string `json:"status" comment:"状态（1-正常 2-异常）"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// ContentArticleGetReq 功能获取请求参数
type ContentArticleGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// ContentArticleDeleteReq 功能删除请求参数
type ContentArticleDeleteReq struct {
	Ids []int64 `json:"ids"`
}
