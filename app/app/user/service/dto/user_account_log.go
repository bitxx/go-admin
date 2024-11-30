package dto

import (
	"github.com/shopspring/decimal"
	commDto "go-admin/app/app/common/dto"
	"go-admin/core/dto"
	"time"
)

type UserAccountLogQueryReq struct {
	dto.Pagination   `search:"-"`
	UserId           int64  `form:"userId"  search:"type:exact;column:user_id;table:app_user_account_log" comment:"用户编号"`
	MoneyType        string `form:"moneyType"  search:"type:exact;column:money_type;table:app_user_account_log" comment:"金额类型 1:余额 "`
	ChangeType       string `form:"changeType"  search:"type:exact;column:change_type;table:app_user_account_log" comment:"帐变类型(1-类型1)"`
	BeginCreatedAt   string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_user_account_log" comment:"创建时间"`
	EndCreatedAt     string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_user_account_log" comment:"创建时间"`
	ShowInfo         bool   `form:"-"  search:"-" comment:"是否明文显示加密信息"`
	commDto.UserJoin `search:"type:inner;on:id:user_id;table:app_user_account_log;join:app_user"`
	UserAccountLogOrder
}

type UserAccountLogOrder struct {
	IdOrder          int64           `form:"idOrder"  search:"type:order;column:id;table:app_user_account_log"`
	UserIdOrder      int64           `form:"userIdOrder"  search:"type:order;column:user_id;table:app_user_account_log"`
	ChangeMoneyOrder decimal.Decimal `form:"changeMoneyOrder"  search:"type:order;column:change_money;table:app_user_account_log"`
	BeforeMoneyOrder decimal.Decimal `form:"beforeMoneyOrder"  search:"type:order;column:before_money;table:app_user_account_log"`
	AfterMoneyOrder  decimal.Decimal `form:"afterMoneyOrder"  search:"type:order;column:after_money;table:app_user_account_log"`
	MoneyTypeOrder   string          `form:"moneyTypeOrder"  search:"type:order;column:money_type;table:app_user_account_log"`
	ChangeTypeOrder  string          `form:"changeTypeOrder"  search:"type:order;column:change_type;table:app_user_account_log"`
	StatusOrder      string          `form:"statusOrder"  search:"type:order;column:status;table:app_user_account_log"`
	CreateByOrder    int64           `form:"createByOrder"  search:"type:order;column:create_by;table:app_user_account_log"`
	CreatedAtOrder   *time.Time      `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user_account_log"`
	UpdateByOrder    int64           `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user_account_log"`
	UpdatedDateOrder *time.Time      `form:"updatedDateOrder"  search:"type:order;column:updated_date;table:app_user_account_log"`
	RemarksOrder     string          `form:"remarksOrder"  search:"type:order;column:remarks;table:app_user_account_log"`
}

func (m *UserAccountLogQueryReq) GetNeedSearch() interface{} {
	return *m
}

// UserAccountLogGetReq 功能获取请求参数
type UserAccountLogGetReq struct {
	Id int64 `uri:"id"`
}
