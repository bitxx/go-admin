package database

import (
	"github.com/jason-wj/logger/logbase"
	"go-admin/common/core"
	toolsConfig "go-admin/common/core/config"
	"go-admin/common/core/pkg"
	"go-admin/common/core/pkg/casbin"
	"gorm.io/driver/mysql"
	"time"

	toolsDB "go-admin/common/core/tools/database"
	. "go-admin/common/core/tools/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Setup 配置数据库
func Setup() {
	for k := range toolsConfig.DatabasesConfig {
		setupSimpleDatabase(k, toolsConfig.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *toolsConfig.Database) {
	logbase.Infof("%s => %s", host, pkg.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					logbase.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, mysql.Open)

	if err != nil {
		logbase.Fatal(pkg.Red(c.Driver+" connect error :"), err)
	} else {
		logbase.Info(pkg.Green(c.Driver + " connect success !"))
	}

	e := mycasbin.Setup(db, "sys_")

	core.Runtime.SetDb(host, db)
	core.Runtime.SetCasbin(host, e)
}
