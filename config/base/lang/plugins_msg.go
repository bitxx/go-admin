package lang

import "go-admin/core/lang"

const (
	HelloWorldCode = 888
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[HelloWorldCode] = "HelloWorld!"
}
