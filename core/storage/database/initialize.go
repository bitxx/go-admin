package database

import (
	"go-admin/core/casbin"
	"go-admin/core/config"
	"go-admin/core/config/database"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"go-admin/core/utils/textutils"
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
	registers := make([]database.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = database.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := database.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
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
		log.Fatal(textutils.Red(c.Driver+" connect error :"), err)
	} else {
		log.Info(textutils.Green(c.Driver + " connect success !"))
	}

	e := mycasbin.Setup(db, "sys_")

	runtime.RuntimeConfig.SetDb(host, db)
	runtime.RuntimeConfig.SetCasbin(host, e)
}
