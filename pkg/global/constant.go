package global

/*
 * 需要和字典匹配
 */

const (
	RouteRootPath = "/admin-api"
	ModelName     = "go-admin"
	LoginLog      = "login_log_queue"
	OperateLog    = "operate_log_queue"
	TrafficKey    = "X-Request-Id"
	LoggerKey     = "_go-admin-logger-request"

	// SysStatusOk 通用-正常
	SysStatusOk    = "1"
	SysStatusNotOk = "2"

	// 角色
	RoleKeyAdmin = "admin" //角色类型，超级管理员

	//数据权限类型
	DataScope1 = "1" //全部数据
	DataScope2 = "2" //自定数据权限
	DataScope3 = "3" //本部门数据权限
	DataScope4 = "4" //本部门及以下数据权限
	DataScope5 = "5" //仅本人数据权限
)

const (
	DBDriverMysql    = "mysql"
	DBDriverPostgres = "postgres"
)

const (
	ModeDev  string = "dev"  //开发模式
	ModeTest string = "test" //测试模式
	ModeProd string = "prod" //生产模式
)
