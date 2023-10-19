package api

import (
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"go-admin/common/global"
	"go-admin/common/utils/strutils"
	"go-admin/config/config"
)

type loggerKey struct{}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logbase.Helper {
	var log *logbase.Helper
	l, ok := c.Get(global.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logbase.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := strutils.GenerateMsgIDFromContext(c)
	return config.LoggerConfig.GetLogger(global.TrafficKey, requestId)
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := strutils.GenerateMsgIDFromContext(c)
	c.Set(global.LoggerKey, config.LoggerConfig.GetLogger(global.TrafficKey, requestId))
}
