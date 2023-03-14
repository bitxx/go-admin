package config

import (
	logger2 "go-admin/common/core/pkg/logger"
)

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

// Setup 设置logger
func (e Logger) Setup() {
	logger2.SetupLogger(
		logger2.WithType(e.Type),
		logger2.WithPath(e.Path),
		logger2.WithLevel(e.Level),
		logger2.WithStdout(e.Stdout),
		logger2.WithCap(e.Cap),
	)
}

var LoggerConfig = new(Logger)
