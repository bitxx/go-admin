package log

import (
	"github.com/bitxx/logger"
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"go-admin/core/global"
	"go-admin/core/utils/strutils"
)

var logHelper *logbase.Helper

type LoggerConf struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

func Init(loggerConf LoggerConf) {
	logHelper = logger.NewLogger(
		logger.WithType(loggerConf.Type),
		logger.WithPath(loggerConf.Path),
		logger.WithLevel(loggerConf.Level),
		logger.WithStdout(loggerConf.Stdout),
		logger.WithCap(loggerConf.Cap),
	)
}

func Info(args ...interface{}) {
	logHelper.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logHelper.Infof(template, args...)
}

func Trace(args ...interface{}) {
	logHelper.Trace(args...)
}

func Tracef(template string, args ...interface{}) {
	logHelper.Tracef(template, args...)
}

func Debug(args ...interface{}) {
	logHelper.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logHelper.Debugf(template, args...)
}

func Warn(args ...interface{}) {
	logHelper.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logHelper.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logHelper.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logHelper.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	logHelper.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logHelper.Fatalf(template, args...)
}

func WithError(err error) *logbase.Helper {
	return logHelper.WithError(err)
}

func WithFields(fields map[string]interface{}) *logbase.Helper {
	return logHelper.WithFields(fields)
}

func LevelForGorm() int {
	return logbase.DefaultLogger.Options().Level.LevelForGorm()
}

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
	return WithFields(map[string]interface{}{global.TrafficKey: requestId})
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := strutils.GenerateMsgIDFromContext(c)
	c.Set(global.LoggerKey, WithFields(map[string]interface{}{global.TrafficKey: requestId}))
}
