package lang

import (
	"go-admin/core/lang"
	"go-admin/core/utils/log"
)

// 多语言翻译 i18n
const (
	// 等级
	AppUserLevelNameAndTypeExistCode = 40000
	AppUserLevelNameEmptyCode        = 40001
	AppUserLevelTypeEmptyCode        = 40002
	AppUserLevelEmptyCode            = 40003
	AppUserLevelHasUsedCode          = 40004

	//用户
	AppUserEmailOrMobileNeedCode   = 40100
	AppUserMobileNeedTitleCode     = 40101
	AppUserEmailFormatErrCode      = 40102
	AppUserMobileFormatErrCode     = 40103
	AppUserMobileEncryptErrCode    = 40104
	AppUserMobileEncryptErrLogCode = 40105
	AppUserEmailEncryptErrCode     = 40106
	AppUserEmailEncryptErrLogCode  = 40107
	AppUserAccountExistLogCode     = 40108
	AppUserMobileExistLogCode      = 40109
	AppUserEmailExistLogCode       = 40110
	AppUserStatusEmptyCode         = 40111
	AppUserActionTypeEmptyCode     = 40112
	AppUserIdEmptyCode             = 40113
	AppUserRefCodeErrLogCode       = 40114
	AppUserRegisterErrCode         = 40115

	//国家区号
	AppUserCountryEmptyCode        = 40116
	AppUserCountryCodeEmptyCode    = 40117
	AppUserCountryStatusEmptyCode  = 40118
	AppUserCountryHasExistCode     = 40119
	AppUserCountryCodeHasExistCode = 40120
)

var (
	MsgInfo = map[int]string{
		AppUserLevelNameAndTypeExistCode: "该等级名称和对应类型已存在!",
		AppUserLevelNameEmptyCode:        "等级名称不得为空",
		AppUserLevelTypeEmptyCode:        "等级类型不得为空",
		AppUserLevelEmptyCode:            "等级不得为空",
		AppUserLevelHasUsedCode:          "等级已被使用",
		AppUserMobileEncryptErrLogCode:   "手机号加密失败：%s",
		AppUserMobileEncryptErrCode:      "手机号加密失败",
		AppUserEmailEncryptErrLogCode:    "邮箱加密失败：%s",
		AppUserEmailEncryptErrCode:       "邮箱加密失败",

		//用户
		AppUserEmailOrMobileNeedCode: "邮箱或手机号至少需要一样",
		AppUserMobileNeedTitleCode:   "手机号和区号需同时输入",
		AppUserEmailFormatErrCode:    "邮箱格式错误",
		AppUserMobileFormatErrCode:   "手机号格式错误",
		AppUserAccountExistLogCode:   "账号：%s 已存在",
		AppUserMobileExistLogCode:    "该手机号已被其他用户使用，不得使用！",
		AppUserEmailExistLogCode:     "该邮箱已被其他用户使用，不得使用！",
		AppUserStatusEmptyCode:       "用户状态不得为空",
		AppUserActionTypeEmptyCode:   "用户行为类型不得为空",
		AppUserIdEmptyCode:           "用户编号",
		AppUserRefCodeErrLogCode:     "推荐吗异常：%s",
		AppUserRegisterErrCode:       "注册异常",

		//国家区号
		AppUserCountryEmptyCode:        "国家名称不得为空",
		AppUserCountryCodeEmptyCode:    "区号不得为空",
		AppUserCountryStatusEmptyCode:  "状态不得为空",
		AppUserCountryHasExistCode:     "该国家名称已存在",
		AppUserCountryCodeHasExistCode: "该国家区号已存在",
	}
)

// Init 初始化
func Init() {
	for k, v := range MsgInfo {
		if lang.MsgInfo[k] == "" {
			lang.MsgInfo[k] = v
		} else {
			log.Fatal("Your plugin lang code %d is used by system or other plugins,please check")
		}
	}
}
