package config

import (
	"fmt"
	loadconfig "github.com/bitxx/load-config"
	"github.com/bitxx/load-config/source"
)

var (
	_cfg *Settings
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	LoggerConfig.GetLogger().Warn("!!! config change and reload")
}

func (e *Settings) Init() {
	e.init()
	LoggerConfig.GetLogger().Warn("!!! config init")
}

func (e *Settings) init() {
	e.Settings.Logger.Setup()
	e.Settings.multiDatabase()
	e.runCallback()
}

// Config 配置集合
type Config struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Auth        *Auth                 `yaml:"auth"`
	Database    *Database             `yaml:"database"`
	Databases   *map[string]*Database `yaml:"databases"`
	Gen         *Gen                  `yaml:"gen"`
	Cache       *Cache                `yaml:"cache"`
	Queue       *Queue                `yaml:"queue"`
	Locker      *Locker               `yaml:"locker"`
}

// 多db改造
func (e *Config) multiDatabase() {
	if len(*e.Databases) == 0 {
		*e.Databases = map[string]*Database{
			"*": e.Database,
		}

	}
}

// Setup 载入配置文件
func Setup(s source.Source,
	fs ...func()) {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Logger:      LoggerConfig,
			Auth:        AuthConfig,
			Database:    DatabaseConfig,
			Databases:   &DatabasesConfig,
			Gen:         GenConfig,
			Cache:       CacheConfig,
			Queue:       QueueConfig,
			Locker:      LockerConfig,
		},
		callbacks: fs,
	}
	var err error
	loadconfig.DefaultConfig, err = loadconfig.NewConfig(
		loadconfig.WithSource(s),
		loadconfig.WithEntity(_cfg),
	)
	if err != nil {
		LoggerConfig.GetLogger().Fatalf(fmt.Sprintf("New config object fail: %s", err.Error()))
	}
	_cfg.Init()
}
