package models

import (
	"encoding/json"
	"errors"
	"github.com/bitxx/logger/logbase"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
	"time"
)

type SysLoginLog struct {
	Id            int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId        int64      `json:"userId" gorm:"size:11;comment:用户编号"`
	Ipaddr        string     `json:"ipaddr" gorm:"size:255;comment:ip地址"`
	LoginLocation string     `json:"loginLocation" gorm:"size:255;comment:归属地"`
	Browser       string     `json:"browser" gorm:"size:255;comment:浏览器"`
	Os            string     `json:"os" gorm:"size:255;comment:系统"`
	Agent         string     `json:"agent" gorm:"size:255;comment:代理"`
	Platform      string     `json:"platform" gorm:"size:255;comment:固件"`
	LoginTime     *time.Time `json:"loginTime" gorm:"comment:登录时间"`
	Remark        string     `json:"remark" gorm:"size:255;comment:备注"`
	Status        string     `json:"status" gorm:"size:1;comment:状态 1-登录 2-退出"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy      int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy      int64      `json:"updateBy" gorm:"index;comment:更新者"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

// SaveLoginLog 从队列中获取登录日志
func SaveLoginLog(message storage.Messager) (err error) {
	//准备db
	prefix := message.GetPrefix()
	db := runtime.RuntimeConfig.GetDbByKey(prefix)
	if db == nil {
		err = errors.New("db not exist")
		logbase.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		return err
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		logbase.Errorf("json Marshal error, %s", err.Error())
		return err
	}
	var l SysLoginLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		logbase.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	err = db.Create(&l).Error
	if err != nil {
		logbase.Errorf("db create error, %s", err.Error())
		return err
	}
	return nil
}
