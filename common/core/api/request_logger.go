package api

import (
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"go-admin/common/core/pkg"
	"go-admin/config/config"
)

type loggerKey struct{}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logbase.Helper {
	var log *logbase.Helper
	l, ok := c.Get(pkg.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logbase.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := pkg.GenerateMsgIDFromContext(c)
	return config.LoggerConfig.GetLogger(pkg.TrafficKey, requestId)
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := pkg.GenerateMsgIDFromContext(c)
	c.Set(pkg.LoggerKey, config.LoggerConfig.GetLogger(pkg.TrafficKey, requestId))
}
