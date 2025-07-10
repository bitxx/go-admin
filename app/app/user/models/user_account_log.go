package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserAccountLog struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId      int64           `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
	ChangeMoney decimal.Decimal `json:"changeMoney" gorm:"column:change_money;type:decimal(10,2);comment:账变金额"`
	BeforeMoney decimal.Decimal `json:"beforeMoney" gorm:"column:before_money;type:decimal(30,18);comment:账变前金额"`
	AfterMoney  decimal.Decimal `json:"afterMoney" gorm:"column:after_money;type:decimal(30,18);comment:账变后金额"`
	MoneyType   string          `json:"moneyType" gorm:"column:money_type;type:char(10);comment:金额类型 1:余额 "`
	ChangeType  string          `json:"changeType" gorm:"column:change_type;type:varchar(30);comment:帐变类型(1-类型1)"`
	Status      string          `json:"status" gorm:"column:status;type:char(1);comment:状态（1正常 2-异常）"`
	CreateBy    int64           `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdateBy    int64           `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	UpdatedAt   *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
	Remarks     string          `json:"remarks" gorm:"column:remarks;type:varchar(500);comment:备注信息"`
	User        *User           `json:"user" gorm:"foreignkey:user_id"`
}

func (UserAccountLog) TableName() string {
	return "app_user_account_log"
}
