package config

import (
	"github.com/bitxx/logger"
	"github.com/bitxx/logger/logbase"
)

var logInner *logbase.Helper

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

func (e Logger) Setup() {
	logInner = logger.NewLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
	)
}

// GetLoggerWithFields 设置logger
func (e Logger) GetLoggerWithFields(fields map[string]interface{}) *logbase.Helper {
	return logInner.WithFields(fields)
}

// GetLogger 获取logger
func (e Logger) GetLogger() *logbase.Helper {
	return logInner
}

var LoggerConfig = new(Logger)
