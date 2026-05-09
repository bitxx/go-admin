package database

import (
	"fmt"
	"go-admin/pkg/config"
	database2 "go-admin/pkg/config/database"
	"go-admin/pkg/runtime"
	"go-admin/pkg/utils/log"
	"go-admin/pkg/utils/textutils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Setup 配置数据库
func Setup() {
	for k := range config.DatabasesConfig {
		setupSimpleDatabase(k, config.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *config.Database) {
	log.Infof("%s => %s", host, textutils.Green(c.Source))
	registers := make([]database2.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = database2.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := database2.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		panic(fmt.Sprintf(c.Driver+" connect error :", err))
	}
	log.Info(textutils.Green(c.Driver + " connect success !"))

	runtime.RuntimeConfig.SetDb(host, db)
}
