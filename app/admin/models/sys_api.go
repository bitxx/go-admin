package models

import (
	"encoding/json"
	"fmt"
	"go-admin/common/runtime"
	"go-admin/common/utils/storage"
	"strings"
	"time"
)

type SysApi struct {
	Id        int        `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle    string     `json:"handle" gorm:"size:128;comment:handle"`
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
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		fmt.Errorf("json Marshal error, %s", err.Error())
		return err
	}

	var l runtime.Routers
	err = json.Unmarshal(rb, &l)
	if err != nil {
		fmt.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	dbList := runtime.RuntimeConfig.GetDb()
	for _, d := range dbList {
		for _, v := range l.List {
			if v.HttpMethod != "HEAD" ||
				strings.Contains(v.RelativePath, "/static/") ||
				strings.Contains(v.RelativePath, "/sys/tables") {
				err := d.Debug().Where(SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
					Attrs(SysApi{Handle: v.Handler}).
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
	return nil
}
