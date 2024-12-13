package models

import (
	"encoding/json"
	"errors"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"go-admin/core/utils/storage"
	"time"
)

type SysOperLog struct {
	Id            int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId        int64      `json:"userId" gorm:"size:11;comment:用户编号"`
	RequestMethod string     `json:"requestMethod" gorm:"size:128;comment:请求方式"`
	OperUrl       string     `json:"operUrl" gorm:"size:255;comment:访问地址"`
	OperIp        string     `json:"operIp" gorm:"size:128;comment:客户端ip"`
	OperLocation  string     `json:"operLocation" gorm:"size:128;comment:访问位置"`
	Status        string     `json:"status" gorm:"size:4;comment:操作状态"`
	OperTime      *time.Time `json:"operTime" gorm:"comment:操作时间"`
	JsonResult    string     `json:"jsonResult" gorm:"size:255;comment:返回数据"`
	Remark        string     `json:"remark" gorm:"size:255;comment:备注"`
	LatencyTime   string     `json:"latencyTime" gorm:"size:128;comment:耗时"`
	UserAgent     string     `json:"userAgent" gorm:"size:255;comment:ua"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy      int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy      int64      `json:"updateBy" gorm:"index;comment:更新者"`
}

func (SysOperLog) TableName() string {
	return "admin_sys_oper_log"
}

// SaveOperLog 从队列中获取操作日志
func SaveOperLog(message storage.Messager) (err error) {
	//准备db
	db := runtime.RuntimeConfig.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var l SysOperLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	// 超出100个字符返回值截断
	if len(l.JsonResult) > 100 {
		l.JsonResult = l.JsonResult[:100]
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	return nil
}
