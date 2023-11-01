package api

import (
	"context"
	"fmt"
	"github.com/bitxx/load-config/source/file"
	"go-admin/app"
	"go-admin/core/config"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
	"go-admin/core/storage/cache"
	"go-admin/core/storage/database"
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
	"go-admin/app/admin/models"
	"go-admin/core/global"
	"go-admin/core/middleware"
)

var (
	configPath string
	apiCheck   bool
	StartCmd   *cobra.Command
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      config.ApplicationConfig.Name + " server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//初始化权限校验
			auth.InitAuth()

			//国际化-初始化底层
			lang.InitLang()

			//国际化-插件支持
			app.AllLang()

			AppRouters = append(AppRouters, app.AllRouter()...)

			return run()
		},
	}

	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
}

func setup() {
	//1. 读取配置
	config.Setup(
		file.NewSource(file.WithPath(configPath)),
		database.Setup,
		cache.Setup,
	)
	//注册监听函数
	queue := runtime.RuntimeConfig.GetMemoryQueue("")
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperLog)
	queue.Register(global.ApiCheck, models.SaveSysApi)
	go queue.Run()
	log.Info(`starting api server...`)
}

func run() error {
	if config.ApplicationConfig.Mode == global.ModeProd.String() {
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

	if apiCheck {
		var routers = runtime.RuntimeConfig.GetRouter()
		q := runtime.RuntimeConfig.GetMemoryQueue("")
		mp := make(map[string]interface{}, 0)
		mp["List"] = routers
		message, err := runtime.RuntimeConfig.GetStreamMessage("", global.ApiCheck, mp)
		if err != nil {
			log.Infof("GetStreamMessage error, %s \n", err.Error())
			//日志报错错误，不中断请求
		} else {
			err = q.Append(message)
			if err != nil {
				log.Infof("Append message error, %s \n", err.Error())
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接，不考虑https，该服务结偶，由专业的转发工具提供，如nginx
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()
	log.Info(textutils.Red(string(global.LogoContent)))
	tip()
	log.Info(textutils.Green("Server run at:"))
	log.Infof("-  Local:   http://localhost:%d/ \r", config.ApplicationConfig.Port)
	log.Infof("-  Network: http://%s:%d/ \r", iputils.GetLocaHost(), config.ApplicationConfig.Port)
	log.Infof("%s Enter Control + C Shutdown Server \r", strutils.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infof("%s Shutdown Server ... \r", strutils.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Error("Server exiting")

	return nil
}

func tip() {
	usageStr := `欢迎使用 ` + textutils.Green(config.ApplicationConfig.Name+" "+config.ApplicationConfig.Version) + ` 可以使用 ` + textutils.Red(`-h`) + ` 查看命令`
	log.Infof("%s", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		h = gin.New()
		runtime.RuntimeConfig.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
	}
	//r.Use(middleware.Metrics())
	r.Use(middleware.Sentinel()).
		Use(middleware.RequestId()).
		Use(log.SetRequestLogger)

	middleware.InitMiddleware(r)

}
