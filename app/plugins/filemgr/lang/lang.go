package lang

import (
	"go-admin/core/lang"
)

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

var (
	MsgInfo = map[int]string{
		AppPlatformEmptyCode:       "请选择一个平台",
		AppVersionEmptyCode:        "请输入版本号",
		AppDownloadTypeEmptyCode:   "请选择下载类型",
		AppDownloadCode:            "下载失败",
		AppPublishStatusCode:       "请设置发布状态",
		AppTypeCode:                "请选择app类型",
		AppCheckFileUploadCode:     "请选择一个app文件上传",
		AppSelectOneFileUploadCode: "每次仅可上传一个文件",
		AppRemarkCode:              "更新内容不得为空",
		AppExistCode:               "该版本已经存在，请重新输入",
		AppOssGenLogCode:           "oss url生成失败：%s",
		AppUploadLogCode:           "app上传失败：%s",
		AppUploadCode:              "app上传失败",
		AppLocalUrlEmptyCode:       "本地Url根地址为空",
		AppOssUploadLogCode:        "oss app上传失败：%s",
		AppUploadSuccessCode:       "app上传成功",
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
