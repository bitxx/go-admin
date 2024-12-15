package lang

import "go-admin/core/lang"

const (
	AppPlatformEmptyCode       = 30000
	AppVersionEmptyCode        = 30001
	AppDownloadTypeEmptyCode   = 30002
	AppDownloadCode            = 30003
	AppPublishStatusCode       = 30004
	AppTypeCode                = 30005
	AppCheckFileUploadCode     = 30006
	AppSelectOneFileUploadCode = 30007
	AppRemarkCode              = 30008
	AppExistCode               = 30009
	AppOssGenLogCode           = 30010
	AppUploadLogCode           = 30011
	AppUploadCode              = 30012
	AppLocalUrlEmptyCode       = 30013
	AppOssUploadLogCode        = 30014
	AppUploadSuccessCode       = 30015
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[AppPlatformEmptyCode] = "请选择一个平台"
	lang.MsgInfo[AppVersionEmptyCode] = "请输入版本号"
	lang.MsgInfo[AppDownloadTypeEmptyCode] = "请选择下载类型"
	lang.MsgInfo[AppDownloadCode] = "下载失败"
	lang.MsgInfo[AppPublishStatusCode] = "请设置发布状态"
	lang.MsgInfo[AppTypeCode] = "请选择app类型"
	lang.MsgInfo[AppCheckFileUploadCode] = "请选择一个app文件上传"
	lang.MsgInfo[AppSelectOneFileUploadCode] = "每次仅可上传一个文件"
	lang.MsgInfo[AppRemarkCode] = "更新内容不得为空"
	lang.MsgInfo[AppExistCode] = "该版本已经存在，请重新输入"
	lang.MsgInfo[AppOssGenLogCode] = "oss url生成失败：%s"
	lang.MsgInfo[AppUploadLogCode] = "app上传失败：%s"
	lang.MsgInfo[AppUploadCode] = "app上传失败"
	lang.MsgInfo[AppLocalUrlEmptyCode] = "本地Url根地址为空"
	lang.MsgInfo[AppOssUploadLogCode] = "oss app上传失败：%s"
	lang.MsgInfo[AppUploadSuccessCode] = "app上传成功"
}
