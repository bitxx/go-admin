package global

/*
 * 需要和字典匹配
 */

const (
	TrafficKey = "X-Request-Id"
	LoggerKey  = "_go-admin-logger-request"

	// SysStatusOk 通用-正常
	SysStatusOk    = "1"
	SysStatusNotOk = "2"
)

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"  //开发模式
	ModeTest Mode = "test" //测试模式
	ModeProd Mode = "prod" //生产模式
)

func (e Mode) String() string {
	return string(e)
}
