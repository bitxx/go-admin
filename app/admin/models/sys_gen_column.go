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
	IsEdit        string     `gorm:"column:is_edit;size:1;" json:"isEdit"`
	IsMust        string     `gorm:"column:is_must;size:1;" json:"isMust"`
	IsQuery       string     `gorm:"column:is_query;size:1;" json:"isQuery"`
	IsList        string     `gorm:"column:is_list;size:1;" json:"isList"`
	QueryType     string     `gorm:"column:query_type;size:128;" json:"queryType"`
	HtmlType      string     `gorm:"column:html_type;size:128;" json:"htmlType"`
	DictType      string     `gorm:"column:dict_type;size:128;" json:"dictType"`
	Sort          int        `gorm:"column:sort;" json:"sort"`
	Remark        string     `gorm:"column:remark;size:255;" json:"remark"`
	CreateBy      int64      `gorm:"column:create_by;size:11;" json:"createBy"`
	UpdateBy      int64      `gorm:"column:update_By;size:11;" json:"updateBy"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysGenColumn) TableName() string {
	return "sys_gen_column"
}

type DBColumn struct {
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TBName                 string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

func (DBColumn) TableName() string {
	return "information_schema.columns"
}
