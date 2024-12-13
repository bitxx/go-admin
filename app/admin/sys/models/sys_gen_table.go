package models

import "time"

type SysGenTable struct {
	Id             int64          `gorm:"primaryKey;autoIncrement" json:"id"`           //表编码
	TBName         string         `gorm:"column:table_name;size:255;" json:"tableName"` //表名称
	TableComment   string         `gorm:"size:255;" json:"tableComment"`                //表备注
	ClassName      string         `gorm:"size:255;" json:"className"`                   //类名
	PackageName    string         `gorm:"size:255;" json:"packageName"`                 //包名
	BusinessName   string         `gorm:"size:255;" json:"businessName"`                //业务名
	ModuleName     string         `gorm:"size:255;" json:"moduleName"`                  //go文件名
	FunctionName   string         `gorm:"size:255;" json:"functionName"`                //功能名称
	FunctionAuthor string         `gorm:"size:255;" json:"functionAuthor"`              //功能作者
	Remark         string         `gorm:"size:255;" json:"remark"`
	CreateBy       int            `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy       int            `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt      *time.Time     `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt"`
	SysGenColumns  []SysGenColumn `json:"sysGenColumns" gorm:"foreignKey:table_id"`
}

func (SysGenTable) TableName() string {
	return "admin_sys_gen_table"
}

type DBTable struct {
	TBName         string     `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string     `gorm:"column:ENGINE" json:"engine"`
	TableRows      string     `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string     `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     *time.Time `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     *time.Time `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string     `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}

func (DBTable) TableName() string {
	return "information_schema.tables"
}
