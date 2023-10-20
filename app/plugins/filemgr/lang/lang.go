package lang

import (
	"go-admin/core/lang"
)

const (
	PluginsAppPlatformEmptyCode       = 30000
	PluginsAppVersionEmptyCode        = 30001
	PluginsAppDownloadTypeEmptyCode   = 30002
	PluginsAppDownloadCode            = 30003
	PluginsAppPublishStatusCode       = 30004
	PluginsAppTypeCode                = 30005
	PluginsAppCheckFileUploadCode     = 30006
	PluginsAppSelectOneFileUploadCode = 30007
	PluginsAppRemarkCode              = 30008
	PluginsAppExistCode               = 30009
	PluginsAppOssGenLogCode           = 30010
	PluginsAppUploadLogCode           = 30011
	PluginsAppUploadCode              = 30012
	PluginsAppLocalUrlEmptyCode       = 30013
	PluginsAppOssUploadLogCode        = 30014
	PluginsAppUploadSuccessCode       = 30015
)

var (
	MsgInfo = map[int]string{
		PluginsAppPlatformEmptyCode:       "请选择一个平台",
		PluginsAppVersionEmptyCode:        "请输入版本号",
		PluginsAppDownloadTypeEmptyCode:   "请选择下载类型",
		PluginsAppDownloadCode:            "下载失败",
		PluginsAppPublishStatusCode:       "请设置发布状态",
		PluginsAppTypeCode:                "请选择app类型",
		PluginsAppCheckFileUploadCode:     "请选择一个app文件上传",
		PluginsAppSelectOneFileUploadCode: "每次仅可上传一个文件",
		PluginsAppRemarkCode:              "更新内容不得为空",
		PluginsAppExistCode:               "该版本已经存在，请重新输入",
		PluginsAppOssGenLogCode:           "oss url生成失败：%s",
		PluginsAppUploadLogCode:           "app上传失败：%s",
		PluginsAppUploadCode:              "app上传失败",
		PluginsAppLocalUrlEmptyCode:       "本地Url根地址",
		PluginsAppOssUploadLogCode:        "oss app上传失败：%s",
		PluginsAppUploadSuccessCode:       "app上传成功",
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
