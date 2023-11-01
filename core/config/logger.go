package config

import (
	"go-admin/core/utils/log"
)

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

func (e Logger) Setup() {

	log.Init(log.LoggerConf{
		Type:      e.Type,
		Path:      e.Path,
		Level:     e.Level,
		Stdout:    e.Stdout,
		EnabledDB: e.EnabledDB,
		Cap:       e.Cap,
	})
}

var LoggerConfig = new(Logger)
