package dto

import (
	"go-admin/common/dto"
	"time"
)

type UserLevelQueryReq struct {
	dto.Pagination `search:"-"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:app_user_level" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:app_user_level" comment:"创建时间"`
	Name           string `form:"name"  search:"type:contains;column:name;table:app_user_level" comment:"等级名称"`
	LevelType      string `form:"levelType"  search:"type:exact;column:level_type;table:app_user_level" comment:"等级类型"`
	Level          int64  `form:"level"  search:"type:exact;column:level;table:app_user_level" comment:"等级"`
	UserLevelOrder
}

type UserLevelOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:app_user_level"`
	NameOrder      string     `form:"nameOrder"  search:"type:order;column:name;table:app_user_level"`
	LevelTypeOrder string     `form:"levelTypeOrder"  search:"type:order;column:level_type;table:app_user_level"`
	LevelOrder     int64      `form:"levelOrder"  search:"type:order;column:level;table:app_user_level"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:app_user_level"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:app_user_level"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:app_user_level"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user_level"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user_level"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:app_user_level"`
}

func (m *UserLevelQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserLevelInsertReq struct {
	Name       string `json:"name" comment:"等级名称"`
	LevelType  string `json:"levelType" comment:"等级类型"`
	Level      int64  `json:"level" comment:"等级"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

type UserLevelUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"等级编号"` // 等级编号
	Name       string `json:"name" comment:"等级名称"`
	LevelType  string `json:"levelType" comment:"等级类型"`
	Level      int64  `json:"level" comment:"等级"`
	CurrUserId int64  `json:"-" comment:"当前登陆用户"`
}

// UserLevelGetReq 功能获取请求参数
type UserLevelGetReq struct {
	Id int64 `uri:"id"`
}

// UserLevelDeleteReq 功能删除请求参数
type UserLevelDeleteReq struct {
	Ids []int64 `json:"ids"`
}
