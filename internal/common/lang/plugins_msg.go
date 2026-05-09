package lang

import "go-admin/pkg/lang"

const (
	HelloWorldCode = 888
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[HelloWorldCode] = "HelloWorld!"
}
