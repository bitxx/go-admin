package lang

import "go-admin/core/lang"

const (
	//600（包含600）以下必须符合http规则，否则接口会异常
	SuccessCode       = 200
	RequestErr        = 400
	AuthErr           = 401
	ForbitErr         = 403
	ServerErr         = 500
	ParamErrCode      = 1001
	OpErrCode         = 1002
	DataDecodeCode    = 1013
	DataDecodeLogCode = 1012
	DataQueryCode     = 1003
	DataQueryLogCode  = 1004
	DataInsertLogCode = 1005
	DataInsertCode    = 1006
	DataNotUpdateCode = 1014
	DataUpdateCode    = 1007
	DataUpdateLogCode = 1008
	DataDeleteCode    = 1009
	DataDeleteLogCode = 1010
	DataNotFoundCode  = 1011
)

func init() {
	//1-基础通用
	lang.MsgInfo[SuccessCode] = "操作成功"
	lang.MsgInfo[RequestErr] = "请求失败"
	lang.MsgInfo[AuthErr] = "状态失效，请重新登录"
	lang.MsgInfo[ForbitErr] = "对不起，您权限不足，操作异常，请联系管理员"
	lang.MsgInfo[ServerErr] = "内部错误"
	lang.MsgInfo[ParamErrCode] = "参数错误"
	lang.MsgInfo[OpErrCode] = "操作异常，请检查"
	lang.MsgInfo[DataDecodeCode] = "数据解析异常"
	lang.MsgInfo[DataDecodeLogCode] = "数据解析异常：%s"
	lang.MsgInfo[DataQueryCode] = "数据查询失败"
	lang.MsgInfo[DataQueryLogCode] = "数据查询失败：%s"
	lang.MsgInfo[DataInsertLogCode] = "数据新增失败：%s"
	lang.MsgInfo[DataInsertCode] = "数据新增失败"
	lang.MsgInfo[DataNotUpdateCode] = "数据未变更"
	lang.MsgInfo[DataUpdateCode] = "数据更新异常"
	lang.MsgInfo[DataUpdateLogCode] = "数据更新异常：%s"
	lang.MsgInfo[DataDeleteCode] = "数据删除失败"
	lang.MsgInfo[DataDeleteLogCode] = "数据删除失败：%s"
	lang.MsgInfo[DataNotFoundCode] = "数据不存在"
}
