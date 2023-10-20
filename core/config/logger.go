package config

import (
	"github.com/bitxx/logger"
	"github.com/bitxx/logger/logbase"
	"strings"
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

// GetLogger 设置logger
func (e Logger) GetLogger(key, requestId string) *logbase.Helper {
	return logInner.WithFields(map[string]interface{}{
		strings.ToLower(key): requestId,
	})
}

var LoggerConfig = new(Logger)
