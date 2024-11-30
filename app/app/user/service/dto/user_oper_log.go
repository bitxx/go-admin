package dto

import (
	commDto "go-admin/app/app/common/dto"
	"go-admin/core/dto"
	"time"
)

type UserOperLogQueryReq struct {
	dto.Pagination   `search:"-"`
	ActionType       string `form:"actionType"  search:"type:exact;column:action_type;table:app_user_oper_log" comment:"用户行为类型"`
	ByType           string `form:"byType"  search:"type:exact;column:by_type;table:app_user_oper_log" comment:"更新用户类型 1-app用户 2-后台用户"`
	UserId           int64  `form:"userId"  search:"type:exact;column:user_id;table:app_user_oper_log" comment:"用户编号"`
	BeginCreatedAt   string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_user_oper_log" comment:"创建时间"`
	EndCreatedAt     string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_user_oper_log" comment:"创建时间"`
	ShowInfo         bool   `form:"-"  search:"-" comment:"是否明文显示加密信息"`
	commDto.UserJoin `search:"type:inner;on:id:user_id;table:app_user_oper_log;join:app_user"`
	UserOperLogOrder
}

type UserOperLogOrder struct {
	ActionTypeOrder string     `form:"actionTypeOrder"  search:"type:order;column:action_type;table:app_user_oper_log"`
	ByTypeOrder     string     `form:"byTypeOrder"  search:"type:order;column:by_type;table:app_user_oper_log"`
	CreateByOrder   int64      `form:"createByOrder"  search:"type:order;column:create_by;table:app_user_oper_log"`
	CreatedAtOrder  *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user_oper_log"`
	IdOrder         int64      `form:"idOrder"  search:"type:order;column:id;table:app_user_oper_log"`
	RemarkOrder     string     `form:"remarkOrder"  search:"type:order;column:remark;table:app_user_oper_log"`
	StatusOrder     string     `form:"statusOrder"  search:"type:order;column:status;table:app_user_oper_log"`
	UpdateByOrder   int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user_oper_log"`
	UpdatedAtOrder  *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:app_user_oper_log"`
	UserIdOrder     int64      `form:"userIdOrder"  search:"type:order;column:user_id;table:app_user_oper_log"`
}

func (m *UserOperLogQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserOperLogInsertReq struct {
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
	ActionType string `json:"actionType" comment:"行为类型"`
	UserId     int64  `json:"userId" comment:"用户编号"`
}

// UserOperLogGetReq 功能获取请求参数
type UserOperLogGetReq struct {
	Id int64 `uri:"id"`
}
