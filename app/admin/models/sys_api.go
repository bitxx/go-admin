package models

import (
	"encoding/json"
	"fmt"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
	"strings"
	"time"
)

var IsSync = false

type SysApi struct {
	Id        int        `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Title     string     `json:"title" gorm:"size:128;comment:标题"`
	Path      string     `json:"path" gorm:"size:128;comment:地址"`
	Action    string     `json:"action" gorm:"size:16;comment:请求类型"`
	ApiType   string     `json:"apiType" gorm:"size:16;comment:接口类型"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
}

func (SysApi) TableName() string {
	return "sys_api"
}

func SaveSysApi(message storage.Messager) (err error) {
	IsSync = true
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		return err
	}

	var l runtime.Routers
	err = json.Unmarshal(rb, &l)
	if err != nil {
		return err
	}
	dbList := runtime.RuntimeConfig.GetDb()
	for _, d := range dbList {
		for _, v := range l.List {
			if v.HttpMethod != "HEAD" ||
				strings.Contains(v.RelativePath, "/static/") ||
				strings.Contains(v.RelativePath, "/sys/tables") {
				err := d.Debug().Where(SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
					FirstOrCreate(&SysApi{}).
					//Update("handle", v.Handler).
					Error
				if err != nil {
					err := fmt.Errorf("Models SaveSysApi error: %s \r\n ", err.Error())
					return err
				}
			}
		}
	}
	IsSync = false
	return nil
}
