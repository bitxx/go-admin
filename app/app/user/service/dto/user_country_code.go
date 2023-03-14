package dto

import (
	"go-admin/common/dto"
	"time"
)

type UserCountryCodeQueryReq struct {
	dto.Pagination `search:"-"`
	Country        string `form:"country"  search:"type:contains;column:country;table:app_user_country_code" comment:"国家地区"`
	CountryInner   string `form:"-"  search:"type:exact;column:country;table:app_user_country_code" comment:"国家地区-内部精确查询"`
	Code           string `form:"code"  search:"type:exact;column:code;table:app_user_country_code" comment:"区号"`
	Status         string `form:"status"  search:"type:exact;column:status;table:app_user_country_code" comment:"状态(1-可用 2-停用)"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:app_user_country_code" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:app_user_country_code" comment:"创建时间"`
	UserCountryCodeOrder
}

type UserCountryCodeOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:app_user_country_code"`
	CountryOrder   string     `form:"countryOrder"  search:"type:order;column:country;table:app_user_country_code"`
	CodeOrder      string     `form:"codeOrder"  search:"type:order;column:code;table:app_user_country_code"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:app_user_country_code"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:app_user_country_code"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:app_user_country_code"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user_country_code"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user_country_code"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:app_user_country_code"`
}

func (m *UserCountryCodeQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserCountryCodeInsertReq struct {
	Country    string `json:"country" comment:"国家地区"`
	Code       string `json:"code" comment:"区号"`
	Status     string `json:"status" comment:"状态(1-可用 2-停用)"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type UserCountryCodeUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"编号"` // 编号
	Country    string `json:"country" comment:"国家地区"`
	Code       string `json:"code" comment:"区号"`
	Status     string `json:"status" comment:"状态(1-可用 2-停用)"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// UserCountryCodeGetReq 功能获取请求参数
type UserCountryCodeGetReq struct {
	Id int64 ` json:"-" uri:"id"`
}

// UserCountryCodeDeleteReq 功能删除请求参数
type UserCountryCodeDeleteReq struct {
	Ids []int64 `json:"ids"`
}
