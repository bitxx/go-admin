package dto

import (
	commDto "go-admin/app/app/common/dto"
	"go-admin/core/dto"
	"time"
)

type UserConfQueryReq struct {
	dto.Pagination   `search:"-"`
	UserId           int64  `form:"userId"  search:"type:exact;column:user_id;table:app_user_conf" comment:"用户编号"`
	CanLogin         string `form:"canLogin"  search:"type:exact;column:can_login;table:app_user_conf" comment:"1-允许登陆；2-不允许登陆"`
	BeginCreatedAt   string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_user_conf" comment:"创建时间"`
	EndCreatedAt     string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_user_conf" comment:"创建时间"`
	ShowInfo         bool   `form:"-"  search:"-" comment:"是否明文显示加密信息"`
	commDto.UserJoin `search:"type:inner;on:id:user_id;table:app_user_conf;join:app_user"`
	UserConfOrder
}

type UserConfOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:app_user_conf"`
	UserIdOrder    int64      `form:"userIdOrder"  search:"type:order;column:user_id;table:app_user_conf"`
	CanLoginOrder  string     `form:"canLoginOrder"  search:"type:order;column:can_login;table:app_user_conf"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:app_user_conf"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:app_user_conf"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:app_user_conf"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user_conf"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user_conf"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:app_user_conf"`
}

func (m *UserConfQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserConfInsertReq struct {
	UserId     int64  `json:"userId" comment:"用户编号"`
	CanLogin   string `json:"canLogin" comment:"1-允许登陆；2-不允许登陆"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type UserConfUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"配置编号"` // 配置编号
	UserId     int64  `json:"userId" comment:"用户编号"`
	CanLogin   string `json:"canLogin" comment:"1-允许登陆；2-不允许登陆"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// UserConfGetReq 功能获取请求参数
type UserConfGetReq struct {
	Id int64 `uri:"id"`
}
