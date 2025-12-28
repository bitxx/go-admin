package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/bitxx/load-config/source/file"
	"go-admin/app"
	mycasbin "go-admin/core/casbin"
	"go-admin/core/config"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
	"go-admin/core/storage/cache"
	"go-admin/core/storage/database"
	"go-admin/core/storage/locker"
	queueSetup "go-admin/core/storage/queue"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
	"go-admin/core/utils/strutils"
	"go-admin/core/utils/textutils"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-admin/app/admin/sys/models"
	"go-admin/core/global"
	"go-admin/core/middleware"
)

var (
	configPath string
	StartCmd   *cobra.Command
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//初始化权限校验
			auth.InitAuth()

			//国际化-初始化底层
			lang.InitLang()

			//国际化-业务

			AppRouters = append(AppRouters, app.AllRouter()...)

			return run()
		},
	}

	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 1. 读取配置
	config.Setup(
		file.NewSource(file.WithPath(configPath)),
		database.Setup,
		cache.Setup,
		queueSetup.Setup,
		locker.Setup,
	)

	// 2.casbin设置
	for host := range config.DatabasesConfig {
		db := runtime.RuntimeConfig.GetDbByKey(host)
		e := mycasbin.Setup(db, "admin_sys_")
		runtime.RuntimeConfig.SetCasbin(host, e)
	}

	// 3. 注册监听函数
	queue := runtime.RuntimeConfig.GetMemoryQueue("")
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperLog)
	go queue.Run()
	log.Info(`starting api server...`)
}

func run() error {
	if config.ApplicationConfig.Mode == global.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: runtime.RuntimeConfig.GetEngine(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverErr := make(chan error, 1)
	go func() {
		// 服务连接，不考虑https，该服务结偶，由专业的转发工具提供，如nginx
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Server listen error: %v", err)
			serverErr <- err
		}
	}()
	log.Info(textutils.Red(string(global.LogoContent)))
	tip()
	log.Info(textutils.Green("Server run at:"))
	log.Infof("-  Local:   http://localhost:%d/ \r", config.ApplicationConfig.Port)
	log.Infof("-  Network: http://%s:%d/ \r", iputils.GetLocaHost(), config.ApplicationConfig.Port)
	log.Infof("%s Enter Control + C Shutdown Server \r", strutils.GetCurrentTimeStr())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
		// 正常关闭流程
		log.Infof("%s Shutdown Server ... \r", strutils.GetCurrentTimeStr())
	case err := <-serverErr:
		// 启动失败，直接返回错误
		log.Errorf("Server failed to start: %v", err)
		return err
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown error: %v", err)
		return err // 返回错误给 RunE
	}
	log.Info("Server exiting")

	return nil
}

func tip() {
	usageStr := `欢迎使用 ` + textutils.Green(config.ApplicationConfig.Name+" "+config.ApplicationConfig.Version) + ` 可以使用 ` + textutils.Red(`--help`) + ` 查看命令`
	log.Infof("%s", usageStr)
}

func initRouter() {
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		h = gin.New()
		runtime.RuntimeConfig.SetEngine(h)
	}
	r, ok := h.(*gin.Engine)
	if !ok {
		panic("not support other engine")
	}
	//r.Use(middleware.Metrics())
	r.Use(middleware.RequestId()).Use(log.SetRequestLogger)

	middleware.InitMiddleware(r)
}
