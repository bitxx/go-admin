package dto

import (
	"github.com/shopspring/decimal"
	commDto "go-admin/internal/app/biz/common/dto"
	"go-admin/pkg/dto"
	"time"
)

type UserQueryReq struct {
	dto.Pagination    `search:"-"`
	BeginCreatedAt    string   `form:"beginCreatedAt" search:"type:gte;column:created_at;table:biz_user" comment:"创建时间"`
	EndCreatedAt      string   `form:"endCreatedAt" search:"type:lte;column:created_at;table:biz_user" comment:"创建时间"`
	LevelId           int64    `form:"levelId"  search:"type:exact;column:level_id;table:biz_user" comment:"用户等级编号"`
	LevelIds          []int64  `form:"levelId"  search:"type:exact;column:level_id;table:biz_user" comment:"用户等级编号"`
	UserName          string   `form:"userName"  search:"type:exact;column:user_name;table:biz_user" comment:"用户昵称"`
	TrueName          string   `form:"trueName"  search:"type:exact;column:true_name;table:biz_user" comment:"真实姓名"`
	Email             string   `form:"email"  search:"type:exact;column:email;table:biz_user" comment:"电子邮箱"`
	MobileTitle       string   `form:"mobileTitle"  search:"type:exact;column:mobile_title;table:biz_user" comment:"国家区号"`
	MobileTitles      []string `form:"-"  search:"type:in;column:mobile_title;table:biz_user" comment:"国家区号列表"`
	Mobile            string   `form:"mobile"  search:"type:exact;column:mobile;table:biz_user" comment:"手机号码"`
	RefCode           string   `form:"-"  search:"type:exact;column:ref_code;table:biz_user" comment:"推荐码"`
	ParentId          int64    `form:"parentId"  search:"type:exact;column:parent_id;table:biz_user" comment:"父级编号"`
	Status            string   `form:"status" search:"type:exact;column:status;table:biz_user" comment:"状态"`
	ShowInfo          bool     `form:"-"  search:"-" comment:"是否明文显示加密信息"`
	commDto.LevelJoin `search:"type:inner;on:id:level_id;table:biz_user;join:biz_user_level"`
	//扩展
	ParentRefCode string `form:"parentRefCode"  search:"-" comment:"上级用户邀请码"`
	UserOrder
}

type UserOrder struct {
	IdOrder          int64           `form:"idOrder"  search:"type:order;column:id;table:biz_user"`
	LevelIdOrder     int64           `form:"levelIdOrder"  search:"type:order;column:level_id;table:biz_user"`
	UserNameOrder    string          `form:"userNameOrder"  search:"type:order;column:user_name;table:biz_user"`
	TrueNameOrder    string          `form:"trueNameOrder"  search:"type:order;column:true_name;table:biz_user"`
	MoneyOrder       decimal.Decimal `form:"moneyOrder"  search:"type:order;column:money;table:biz_user"`
	EmailOrder       string          `form:"emailOrder"  search:"type:order;column:email;table:biz_user"`
	MobileTitleOrder string          `form:"mobileTitleOrder"  search:"type:order;column:mobile_title;table:biz_user"`
	MobileOrder      string          `form:"mobileOrder"  search:"type:order;column:mobile;table:biz_user"`
	AvatarOrder      string          `form:"avatarOrder"  search:"type:order;column:avatar;table:biz_user"`
	PayPwdOrder      string          `form:"payPwdOrder"  search:"type:order;column:pay_pwd;table:biz_user"`
	PwdOrder         string          `form:"pwdOrder"  search:"type:order;column:pwd;table:biz_user"`
	RefCodeOrder     string          `form:"refCodeOrder"  search:"type:order;column:ref_code;table:biz_user"`
	ParentIdOrder    int64           `form:"parentIdOrder"  search:"type:order;column:parent_id;table:biz_user"`
	ParentIdsOrder   string          `form:"parentIdsOrder"  search:"type:order;column:parent_ids;table:biz_user"`
	TreeSortOrder    decimal.Decimal `form:"treeSortOrder"  search:"type:order;column:tree_sort;table:biz_user"`
	TreeSortsOrder   string          `form:"treeSortsOrder"  search:"type:order;column:tree_sorts;table:biz_user"`
	TreeLeafOrder    string          `form:"treeLeafOrder"  search:"type:order;column:tree_leaf;table:biz_user"`
	TreeLevelOrder   int64           `form:"treeLevelOrder"  search:"type:order;column:tree_level;table:biz_user"`
	StatusOrder      string          `form:"statusOrder"  search:"type:order;column:status;table:biz_user"`
	RemarkOrder      string          `form:"remarkOrder"  search:"type:order;column:remark;table:biz_user"`
	CreateByOrder    int64           `form:"createByOrder"  search:"type:order;column:create_by;table:biz_user"`
	UpdateByOrder    int64           `form:"updateByOrder"  search:"type:order;column:update_by;table:biz_user"`
	CreatedAtOrder   *time.Time      `form:"createdAtOrder"  search:"type:order;column:created_at;table:biz_user"`
	UpdatedAtOrder   *time.Time      `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:biz_user"`
}

func (m *UserQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserInsertReq struct {
	LevelId     int64           `json:"levelId" comment:"用户等级编号"`
	UserName    string          `json:"userName" comment:"用户昵称"`
	TrueName    string          `json:"trueName" comment:"真实姓名"`
	Money       decimal.Decimal `json:"money" comment:"余额"`
	Email       string          `json:"email" comment:"电子邮箱"`
	MobileTitle string          `json:"mobileTitle" comment:"用户手机号国家前缀"`
	Mobile      string          `json:"mobile" comment:"手机号码"`
	CurrUserId  int64           `json:"-" comment:"当前登陆用户"`
	RefCode     string          `json:"refCode" comment:"邀请码"`
	Emails      string          `json:"emails" comment:"邮箱集合"`
	Mobiles     string          `json:"mobiles" comment:"手机集合"`
}

type UserUpdateReq struct {
	Id          int64           `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	LevelId     int64           `json:"levelId" comment:"用户等级编号"`
	UserName    string          `json:"userName" comment:"用户昵称"`
	TrueName    string          `json:"trueName" comment:"真实姓名"`
	Money       decimal.Decimal `json:"money" comment:"余额"`
	Email       string          `json:"email" comment:"电子邮箱"`
	MobileTitle string          `json:"mobileTitle" comment:"用户手机号国家前缀"`
	Mobile      string          `json:"mobile" comment:"手机号码"`
	CurrUserId  int64           `json:"-" comment:"当前登陆用户"`
}

type UserStatusUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"用户ID"` // 用户ID
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:""`
}

// UserGetReq 功能获取请求参数
type UserGetReq struct {
	Id int64 `uri:"id"`
}
