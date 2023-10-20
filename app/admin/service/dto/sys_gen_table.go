package dto

import (
	"go-admin/core/dto"
)

type SysGenTableQueryReq struct {
	dto.Pagination `search:"-"`
	TableName      string   `form:"tableName" search:"type:contains;column:table_name;table:sys_gen_table" comment:"表名"`
	TableNames     []string `form:"-" search:"type:in;column:table_name;table:sys_gen_table" comment:"表名"`
	TableComment   string   `form:"tableComment" search:"type:icontains;column:table_comment;table:sys_gen_table" comment:"表别名"`
	BeginTime      string   `form:"beginTime" search:"type:gte;column:ctime;table:sys_gen_table" comment:"创建时间"`
	EndTime        string   `form:"endTime" search:"type:lte;column:ctime;table:sys_gen_table" comment:"创建时间"`
	SysGenTableOrder
}

type SysGenTableOrder struct {
	TableName      string `search:"type:order;column:table_name;table:sys_gen_table" form:"tableName"`
	TableComment   string `search:"type:order;column:table_comment;table:sys_gen_table" form:"tableComment"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_gen_table" form:"createdAtOrder"`
}

func (m *SysGenTableQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysGenTableInsertReq struct {
	DbTableNames []string `json:"dbTableNames" comment:"系统表名称集合"`
	CurrUserId   int64    `json:"-" comment:""`
}

type SysGenTableUpdateReq struct {
	Id             int64                   `json:"-" uri:"id" comment:"主键编码"` // 主键编码
	CurrUserId     int64                   `json:"-" comment:""`
	TableComment   string                  `json:"tableComment" comment:"表备注"` //表备注
	ClassName      string                  `json:"className" comment:"类名"`
	PackageName    string                  `json:"packageName" comment:"包名"`
	BusinessName   string                  `json:"businessName" comment:"业务名"`
	ModuleName     string                  `json:"moduleName" comment:"go文件名"`
	FunctionName   string                  `json:"functionName" comment:"功能名称"`
	FunctionAuthor string                  `json:"functionAuthor" comment:"功能作者"`
	Remark         string                  `json:"remark" comment:"备注"`
	Columns        []SysGenColumnUpdateReq `json:"sysGenColumns"`
}

type SysGenTableGetReq struct {
	Id         int64 `uri:"id" json:"-"`
	CurrUserId int64 `json:"-" comment:""`
}

// SysGenTableDeleteReq 功能删除请求参数
type SysGenTableDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type SysGenTableGenCodeReq struct {
	Id         int64  `uri:"id" json:"-"`
	IsDownload string `json:"-" comment:"是否下载zip文件"`
}

// ------------------ DBTable ---------------
type DBTableQueryReq struct {
	dto.Pagination `search:"-"`
	TableName      string `form:"tableName" search:"type:contains;column:TABLE_NAME;table:tables"`
	TableComment   string `form:"tableComment" search:"type:contains;column:TABLE_COMMENT;table:tables"`
}

func (m *DBTableQueryReq) GetNeedSearch() interface{} {
	return *m
}

type DBTableResp struct {
	TableName    string `json:"tableName" comment:"表名称"`
	TableComment string `json:"tableComment" comment:"表描述"`
	CreatedAt    string `json:"createdAt" comment:"创建时间"`
}

type TemplateResp struct {
	Name    string `json:"name" comment:"标记"`
	Path    string `json:"path" comment:"模板路径"`
	Content string `json:"content" comment:"解析后的文本内容"`
}
