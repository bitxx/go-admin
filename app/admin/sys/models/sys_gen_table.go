package models

import (
	"go-admin/core/config"
	"go-admin/core/global"
	"time"
)

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
	TBName         string     `gorm:"column:table_name" json:"tableName"`
	Engine         string     `gorm:"column:engine" json:"engine"`
	TableRows      string     `gorm:"column:table_rows" json:"tableRows"`
	TableCollation string     `gorm:"column:table_collation" json:"tableCollation"`
	CreateTime     *time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     *time.Time `gorm:"column:update_time" json:"updateTime"`
	TableComment   string     `gorm:"column:table_comment" json:"tableComment"`
}

func (DBTable) TableName() string {
	if config.DatabaseConfig.Driver == global.DBDriverPostgres {
		return "pg_catalog.pg_tables"
	}
	return "information_schema.tables"
}
