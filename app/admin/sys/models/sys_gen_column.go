package models

import "time"

type SysGenColumn struct {
	Id            int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	TableId       int64      `gorm:"column:table_id;size:11;" json:"tableId"`
	ColumnName    string     `gorm:"size:128;" json:"columnName"`
	ColumnComment string     `gorm:"column:column_comment;size:128;" json:"columnComment"`
	ColumnType    string     `gorm:"column:column_type;size:128;" json:"columnType"`
	GoType        string     `gorm:"column:go_type;size:128;" json:"goType"`
	GoField       string     `gorm:"column:go_field;size:128;" json:"goField"`
	JsonField     string     `gorm:"column:json_field;size:128;" json:"jsonField"`
	IsPk          string     `gorm:"column:is_pk;size:1;" json:"isPk"`
	IsRequired    string     `gorm:"column:is_required;size:1;" json:"isRequired"`
	IsQuery       string     `gorm:"column:is_query;size:1;" json:"isQuery"`
	IsList        string     `gorm:"column:is_list;size:1;" json:"isList"`
	QueryType     string     `gorm:"column:query_type;size:128;" json:"queryType"`
	HtmlType      string     `gorm:"column:html_type;size:128;" json:"htmlType"`
	DictType      string     `gorm:"column:dict_type;size:128;" json:"dictType"`
	Sort          int        `gorm:"column:sort;" json:"sort"`
	Remark        string     `gorm:"column:remark;size:255;" json:"remark"`
	CreateBy      int64      `gorm:"column:create_by;size:11;" json:"createBy"`
	UpdateBy      int64      `gorm:"column:update_by;size:11;" json:"updateBy"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysGenColumn) TableName() string {
	return "admin_sys_gen_column"
}

type DBColumn struct {
	TableSchema            string `gorm:"column:table_schema" json:"tableSchema"`
	TBName                 string `gorm:"column:table_name" json:"tableName"`
	ColumnName             string `gorm:"column:column_name" json:"columnName"`
	ColumnDefault          string `gorm:"column:column_default" json:"columnDefault"`
	IsNullable             string `gorm:"column:is_nullable" json:"isNullable"`
	DataType               string `gorm:"column:data_type" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:character_maximum_length" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:character_set_name" json:"characterSetName"`
	ColumnType             string `gorm:"column:column_type" json:"columnType"`
	ColumnKey              string `gorm:"column:column_key" json:"columnKey"`
	Extra                  string `gorm:"column:extra" json:"extra"`
	ColumnComment          string `gorm:"column:column_comment" json:"columnComment"`
}

func (DBColumn) TableName() string {
	return "information_schema.columns"
}
