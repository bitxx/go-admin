package api

import (
	"context"
	"fmt"
	"github.com/unrolled/secure"
	"go-admin/app"
	"go-admin/common/core"
	"go-admin/common/core/api"
	"go-admin/common/core/config"
	"go-admin/common/core/config/sdk/source/file"
	"go-admin/common/core/pkg"
	"go-admin/common/core/runtime"
	"go-admin/common/middleware/auth"
	"go-admin/common/storage/cache"
	"go-admin/common/storage/database"
	"go-admin/common/utils/i18n"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-admin/app/admin/models"
	"go-admin/common/global"
	"go-admin/common/middleware"
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
		Example:      "admin-api server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//初始化权限校验
			auth.InitAuth()

			//国际化-初始化底层
			i18n.InitLang()

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
	queue := core.Runtime.GetMemoryQueue("")
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperLog)
	queue.Register(global.ApiCheck, models.SaveSysApi)
	go queue.Run()

	usageStr := `starting api server...`
	log.Println(usageStr)
}

func run() error {
	if config.ApplicationConfig.Mode == pkg.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: core.Runtime.GetEngine(),
	}

	if apiCheck {
		var routers = core.Runtime.GetRouter()
		q := core.Runtime.GetMemoryQueue("")
		mp := make(map[string]interface{}, 0)
		mp["List"] = routers
		message, err := core.Runtime.GetStreamMessage("", global.ApiCheck, mp)
		if err != nil {
			log.Printf("GetStreamMessage error, %s \n", err.Error())
			//日志报错错误，不中断请求
		} else {
			err = q.Append(message)
			if err != nil {
				log.Printf("Append message error, %s \n", err.Error())
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接
		if config.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	fmt.Println(pkg.Red(string(global.LogoContent)))
	tip()
	fmt.Println(pkg.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", pkg.GetLocaHonst(), config.ApplicationConfig.Port)

	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return nil
}

var Router runtime.Router

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-admin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := core.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		core.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	if config.SslConfig.Enable {
		r.Use(TlsHandler())
	}
	//r.Use(middleware.Metrics())
	r.Use(middleware.Sentinel()).
		Use(middleware.RequestId(pkg.TrafficKey)).
		Use(api.SetRequestLogger)

	middleware.InitMiddleware(r)

}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.SslConfig.Domain,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
