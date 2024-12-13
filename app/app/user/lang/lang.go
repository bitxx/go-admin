package lang

import (
	"go-admin/core/lang"
	"go-admin/core/utils/log"
)

// 多语言翻译 i18n
const (
	// 等级
	UserLevelNameAndTypeExistCode = 40000
	UserLevelNameEmptyCode        = 40001
	UserLevelTypeEmptyCode        = 40002
	UserLevelEmptyCode            = 40003
	UserLevelHasUsedCode          = 40004

	//用户
	UserEmailOrMobileNeedCode   = 40100
	UserMobileNeedTitleCode     = 40101
	UserEmailFormatErrCode      = 40102
	UserMobileFormatErrCode     = 40103
	UserMobileEncryptErrCode    = 40104
	UserMobileEncryptErrLogCode = 40105
	UserEmailEncryptErrCode     = 40106
	UserEmailEncryptErrLogCode  = 40107
	UserAccountExistLogCode     = 40108
	UserMobileExistLogCode      = 40109
	UserEmailExistLogCode       = 40110
	UserStatusEmptyCode         = 40111
	UserActionTypeEmptyCode     = 40112
	UserIdEmptyCode             = 40113
	UserRefCodeErrLogCode       = 40114
	UserRegisterErrCode         = 40115

	//国家区号
	UserCountryEmptyCode        = 40116
	UserCountryCodeEmptyCode    = 40117
	UserCountryStatusEmptyCode  = 40118
	UserCountryHasExistCode     = 40119
	UserCountryCodeHasExistCode = 40120
)

var (
	MsgInfo = map[int]string{
		UserLevelNameAndTypeExistCode: "该等级名称和对应类型已存在!",
		UserLevelNameEmptyCode:        "等级名称不得为空",
		UserLevelTypeEmptyCode:        "等级类型不得为空",
		UserLevelEmptyCode:            "等级不得为空",
		UserLevelHasUsedCode:          "等级已被使用",
		UserMobileEncryptErrLogCode:   "手机号加密失败：%s",
		UserMobileEncryptErrCode:      "手机号加密失败",
		UserEmailEncryptErrLogCode:    "邮箱加密失败：%s",
		UserEmailEncryptErrCode:       "邮箱加密失败",

		//用户
		UserEmailOrMobileNeedCode: "邮箱或手机号至少需要一样",
		UserMobileNeedTitleCode:   "手机号和区号需同时输入",
		UserEmailFormatErrCode:    "邮箱格式错误",
		UserMobileFormatErrCode:   "手机号格式错误",
		UserAccountExistLogCode:   "账号：%s 已存在",
		UserMobileExistLogCode:    "该手机号已被其他用户使用，不得使用！",
		UserEmailExistLogCode:     "该邮箱已被其他用户使用，不得使用！",
		UserStatusEmptyCode:       "用户状态不得为空",
		UserActionTypeEmptyCode:   "用户行为类型不得为空",
		UserIdEmptyCode:           "用户编号",
		UserRefCodeErrLogCode:     "推荐吗异常：%s",
		UserRegisterErrCode:       "注册异常",

		//国家区号
		UserCountryEmptyCode:        "国家名称不得为空",
		UserCountryCodeEmptyCode:    "区号不得为空",
		UserCountryStatusEmptyCode:  "状态不得为空",
		UserCountryHasExistCode:     "该国家名称已存在",
		UserCountryCodeHasExistCode: "该国家区号已存在",
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
