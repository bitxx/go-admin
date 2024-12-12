package global

/*
 * 需要和字典匹配
 */

const (
	RouteRootPath = "/admin-api"
	ModelName     = "go-admin"
	LoginLog      = "login_log_queue"
	OperateLog    = "operate_log_queue"
	ApiCheck      = "api_check_queue"
	TrafficKey    = "X-Request-Id"
	LoggerKey     = "_go-admin-logger-request"

	// SysStatusOk 通用-正常
	SysStatusOk    = "1"
	SysStatusNotOk = "2"
)

const (
	GenTypeVue   string = "vue"   //开发模式
	GenTypeReact string = "react" //测试模式
)

const (
	ModeDev  string = "dev"  //开发模式
	ModeTest string = "test" //测试模式
	ModeProd string = "prod" //生产模式
)
