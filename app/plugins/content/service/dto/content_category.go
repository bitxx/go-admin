package dto

import (
	"go-admin/core/dto"
	"time"
)

type ContentCategoryQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int64  `form:"id"  search:"type:exact;column:id;table:plugins_content_category" comment:"编号"`
	Name           string `form:"name"  search:"type:contains;column:name;table:plugins_content_category" comment:"分类名称"`
	NameInner      string `form:"-"  search:"type:exact;column:name;table:plugins_content_category" comment:"分类名称，跟Name区分，数据库精确查询"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:plugins_content_category" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:plugins_content_category" comment:"创建时间"`
	ContentCategoryOrder
}

type ContentCategoryOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:plugins_content_category"`
	NameOrder      string     `form:"nameOrder"  search:"type:order;column:name;table:plugins_content_category"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:plugins_content_category"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:plugins_content_category"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:plugins_content_category"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:plugins_content_category"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:plugins_content_category"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:plugins_content_category"`
}

func (m *ContentCategoryQueryReq) GetNeedSearch() interface{} {
	return *m
}

type ContentCategoryInsertReq struct {
	Name       string `json:"name" comment:"分类名称"`
	Remark     string `json:"remark" comment:"备注信息"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type ContentCategoryUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"分类编号"` // 分类编号
	Name       string `json:"name" comment:"分类名称"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// ContentCategoryGetReq 功能获取请求参数
type ContentCategoryGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// ContentCategoryDeleteReq 功能删除请求参数
type ContentCategoryDeleteReq struct {
	Ids []int64 `json:"ids"`
}
