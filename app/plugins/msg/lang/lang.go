package lang

import "go-admin/config/lang"

// 多语言翻译 i18n
const (
	//编码不得和整个项目中的别的模块有重复
	HelloWorldCode = 888
)

var (
	MsgInfo = map[int]string{
		HelloWorldCode: "HelloWorld!",
	}
)

// Init 初始化
func Init() {
	for k, v := range MsgInfo {
		if lang.MsgInfo[k] == "" {
			lang.MsgInfo[k] = v
		} else {
			panic("Your plugin lang code %d is used by system or other plugins,please check")
		}
	}
}
