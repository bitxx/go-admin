package dto

import (
	"go-admin/core/dto"
)

type SysGenColumnQueryReq struct {
	TableId        int64 `form:"tableId" search:"type:exact;column:table_id;table:admin_sys_gen_column" comment:"表id"`
	dto.Pagination `search:"-"`
	SysGenColumnOrder
}

type SysGenColumnOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:admin_sys_gen_column" form:"createdAtOrder"`
}

func (m *SysGenColumnQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysGenColumnInsertReq struct {
	CurrUserId int64 `json:"-" comment:""`
}

type SysGenColumnUpdateReq struct {
	Id            int64  `json:"id" uri:"id" comment:"主键编码"` // 主键编码
	CurrUserId    int64  `json:"-" comment:""`
	ColumnName    string `json:"columnName"`
	ColumnComment string `json:"columnComment"`
	GoType        string `json:"goType"`
	GoField       string `json:"goField"`
	JsonField     string `json:"jsonField"`
	IsRequired    string `json:"isRequired"`
	IsQuery       string `json:"isQuery"`
	IsList        string `json:"isList"`
	QueryType     string `json:"queryType"`
	HtmlType      string `json:"htmlType"`
	DictType      string `json:"dictType"`
}

type SysGenColumnGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// SysGenTableDeleteReq 功能删除请求参数
type SysGenColumnDeleteReq struct {
	Ids      []int64 `json:"ids"`
	TableIds []int64 `json:"tableIds"`
}

// ------------------ DBTable ---------------
type DBColumnQueryReq struct {
	dto.Pagination `search:"-"`
}

func (m *DBColumnQueryReq) GetNeedSearch() interface{} {
	return *m
}
