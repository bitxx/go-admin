package dto

import (
	"github.com/shopspring/decimal"
	"go-admin/core/dto"
)

type SysDictDataQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int64  `form:"id" search:"type:exact;column:id;table:admin_sys_dict_data" comment:""`
	DictLabel      string `form:"dictLabel" search:"type:contains;column:dict_label;table:admin_sys_dict_data" comment:""`
	DictValue      string `form:"dictValue" search:"type:leftcontains;column:dict_value;table:admin_sys_dict_data" comment:""`
	DictType       string `form:"dictType" search:"type:contains;column:dict_type;table:admin_sys_dict_data" comment:""`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:admin_sys_dict_data" comment:"创建时间"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:admin_sys_dict_data" comment:"创建时间"`
}

func (m *SysDictDataQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysDictDataInsertReq struct {
	DictSort   int    `json:"dictSort" comment:""`
	DictLabel  string `json:"dictLabel" comment:""`
	DictValue  string `json:"dictValue" comment:""`
	DictType   string `json:"dictType" comment:""`
	Remark     string `json:"remark" comment:""`
	CurrUserId int64  `json:"-" comment:""`
}

type SysDictDataUpdateReq struct {
	Id         int64           `uri:"id" json:"-" comment:""`
	DictSort   decimal.Decimal `json:"dictSort" comment:""`
	DictLabel  string          `json:"dictLabel" comment:""`
	DictValue  string          `json:"dictValue" comment:""`
	DictType   string          `json:"dictType" comment:""`
	Remark     string          `json:"remark" comment:""`
	CurrUserId int64           `json:"-" comment:""`
}

type SysDictDataGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

type SysDictDataDeleteReq struct {
	Ids []int64 `json:"ids"`
}
