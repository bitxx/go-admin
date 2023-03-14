package service

import (
	"go-admin/common/core/logger"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Helper
	Lang  string //语言 en 英文 zh-cn中文
}
