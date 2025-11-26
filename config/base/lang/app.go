package lang

import "go-admin/core/lang"

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

	CountryChina       = 40121
	CountryChinaTaiwan = 40122
	CountryChinaMacao  = 40123
	CountryHongkong    = 40124
	CountrySingapore   = 40125
	CountryCanada      = 40126
	CountryKorea       = 40127
	CountryJapan       = 40128
	CountryThailand    = 40129
	CountryBurma       = 40130
	CountryLaos        = 40131
	CountryAustralia   = 40132
	CountryRussia      = 40133
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[UserLevelNameAndTypeExistCode] = "该等级名称和对应类型已存在!"
	lang.MsgInfo[UserLevelNameEmptyCode] = "等级名称不得为空"
	lang.MsgInfo[UserLevelTypeEmptyCode] = "等级类型不得为空"
	lang.MsgInfo[UserLevelEmptyCode] = "等级不得为空"
	lang.MsgInfo[UserLevelHasUsedCode] = "等级已被使用"
	lang.MsgInfo[UserMobileEncryptErrLogCode] = "手机号加密失败：%s"
	lang.MsgInfo[UserMobileEncryptErrCode] = "手机号加密失败"
	lang.MsgInfo[UserEmailEncryptErrLogCode] = "邮箱加密失败：%s"
	lang.MsgInfo[UserEmailEncryptErrCode] = "邮箱加密失败"

	//用户
	lang.MsgInfo[UserEmailOrMobileNeedCode] = "邮箱或手机号至少需要一样"
	lang.MsgInfo[UserMobileNeedTitleCode] = "手机号和区号需同时输入"
	lang.MsgInfo[UserEmailFormatErrCode] = "邮箱格式错误"
	lang.MsgInfo[UserMobileFormatErrCode] = "手机号格式错误"
	lang.MsgInfo[UserAccountExistLogCode] = "账号：%s 已存在"
	lang.MsgInfo[UserMobileExistLogCode] = "该手机号已被其他用户使用，不得使用！"
	lang.MsgInfo[UserEmailExistLogCode] = "该邮箱已被其他用户使用，不得使用！"
	lang.MsgInfo[UserStatusEmptyCode] = "用户状态不得为空"
	lang.MsgInfo[UserActionTypeEmptyCode] = "用户行为类型不得为空"
	lang.MsgInfo[UserIdEmptyCode] = "用户编号"
	lang.MsgInfo[UserRefCodeErrLogCode] = "推荐吗异常：%s"
	lang.MsgInfo[UserRegisterErrCode] = "注册异常"

	//国家区号
	lang.MsgInfo[UserCountryEmptyCode] = "国家名称不得为空"
	lang.MsgInfo[UserCountryCodeEmptyCode] = "区号不得为空"
	lang.MsgInfo[UserCountryStatusEmptyCode] = "状态不得为空"
	lang.MsgInfo[UserCountryHasExistCode] = "该国家名称已存在"
	lang.MsgInfo[UserCountryCodeHasExistCode] = "该国家区号已存在"
	//2-国家
	lang.MsgInfo[CountryChina] = "中国大陆"
	lang.MsgInfo[CountryChinaTaiwan] = "中国台湾"
	lang.MsgInfo[CountryChinaMacao] = "中国澳门"
	lang.MsgInfo[CountryHongkong] = "中国香港"
	lang.MsgInfo[CountrySingapore] = "新加坡"
	lang.MsgInfo[CountryCanada] = "加拿大"
	lang.MsgInfo[CountryKorea] = "韩国"
	lang.MsgInfo[CountryJapan] = "日本"
	lang.MsgInfo[CountryThailand] = "泰国"
	lang.MsgInfo[CountryBurma] = "缅甸"
	lang.MsgInfo[CountryLaos] = "老挝"
	lang.MsgInfo[CountryAustralia] = "澳大利亚"
	lang.MsgInfo[CountryRussia] = "俄罗斯"
}
