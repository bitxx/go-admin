package service

import (
	"github.com/jason-wj/logger/logbase"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logbase.Helper
	Lang  string //语言 en 英文 zh-cn中文
}
