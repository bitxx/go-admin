// Package app
// @Description: 多语言消息管理，根据业务需要新增消息内容，标识内容勿操作
package lang

import (
	"errors"
	"fmt"
	"github.com/jason-wj/logger/logbase"
	"go-admin/common/utils/i18n"
	"strings"
)

const (
	//为避免多语言异常，如无特殊必要，该部分勿修改
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

	CountryChina       = 9500
	CountryChinaTaiwan = 9501
	CountryChinaMacao  = 9502
	CountryHongkong    = 9503
	CountrySingapore   = 9504
	CountryCanada      = 9505
	CountryKorea       = 9506
	CountryJapan       = 9507
	CountryThailand    = 9508
	CountryBurma       = 9509
	CountryLaos        = 9510
	CountryAustralia   = 9511
	CountryRussia      = 9512
)

var (
	MsgInfo = map[int]string{
		//为避免多语言异常，如无特殊必要，该部分勿修改
		//1-基础通用
		SuccessCode:       "操作成功",
		RequestErr:        "请求失败",
		AuthErr:           "状态失效，请重新登录",
		ForbitErr:         "对不起，您没有该接口访问权限，请联系管理员",
		ServerErr:         "内部错误",
		ParamErrCode:      "参数错误",
		OpErrCode:         "操作异常，请检查",
		DataDecodeCode:    "数据解析异常",
		DataDecodeLogCode: "数据解析异常：%s",
		DataQueryCode:     "数据查询失败",
		DataQueryLogCode:  "数据查询失败：%s",
		DataInsertLogCode: "数据新增失败：%s",
		DataInsertCode:    "数据新增失败",
		DataNotUpdateCode: "数据未变更",
		DataUpdateCode:    "数据更新异常",
		DataUpdateLogCode: "数据更新异常：%s",
		DataDeleteCode:    "数据删除失败",
		DataDeleteLogCode: "数据删除失败：%s",
		DataNotFoundCode:  "数据不存在",

		//2-国家
		CountryChina:       "中国大陆",
		CountryChinaTaiwan: "中国台湾",
		CountryChinaMacao:  "中国澳门",
		CountryHongkong:    "中国香港",
		CountrySingapore:   "新加坡",
		CountryCanada:      "加拿大",
		CountryKorea:       "韩国",
		CountryJapan:       "日本",
		CountryThailand:    "泰国",
		CountryBurma:       "缅甸",
		CountryLaos:        "老挝",
		CountryAustralia:   "澳大利亚",
		CountryRussia:      "俄罗斯",

		//业务扩展-以下可修改
	}
)

// MsgByCode
// @Description: i18n
// @param errCode
// @param lang
// @return string
func MsgByCode(errCode int, lang string) string {
	switch lang {
	case "en":
		return i18n.EnI18nClient.T(MsgInfo[errCode])
	default:
		return MsgInfo[errCode]
	}
}

// MsgByValue
// @Description: 直接根据值返回对应语言
// @param value
// @param lang
// @return string
func MsgByValue(value string, lang string) string {
	switch lang {
	case "en":
		return i18n.EnI18nClient.T(value)
	default:
		return value
	}
}

// MsgErr
// @Description: 获取error
// @param errCode
// @param lang
// @return error
func MsgErr(errCode int, lang string) error {
	return errors.New(MsgByCode(errCode, lang))
}

// MsgErrf
// @Description:
// @param errCode
// @param lang
// @param f
// @return error
func MsgErrf(errCode int, lang string, f ...interface{}) error {
	return errors.New(fmt.Sprintf(MsgByCode(errCode, lang), f))
}

// MsgLogErrf
// @Description: 带有参数，有些底层消息不应当被使用者感知，该类消息记录在日志中，并返回应用层可理解的消息
// @param log  用于记录日志，
// @param errCodeReplace 最终需要给应用层返回的消息，这里传入消息码。若errCodeReplace=errCode或errCodeReplace<=0，则返回真实消息
// @param errCode  真实消息码
// @param lang  语言
// @param f
// @return error
func MsgLogErrf(log *logbase.Helper, lang string, errCodeReplace, errCode int, f ...interface{}) error {
	err := MsgErrf(errCode, lang, f)
	log.Error(err)
	if errCodeReplace <= 0 || errCodeReplace == errCode {
		return err
	}
	return MsgErr(errCodeReplace, lang)
}

// MsgLogErr
// @Description: 无参数，有些底层消息不应当被使用者感知，该类消息记录在日志中，并返回应用层可理解的消息
// @param log
// @param lang
// @param errCodeReplace
// @param errCode
// @return error
func MsgLogErr(log *logbase.Helper, lang string, errCodeReplace, errCode int) error {
	err := MsgErr(errCode, lang)
	log.Error(err)
	if errCodeReplace <= 0 || errCodeReplace == errCode {
		return err
	}
	return MsgErr(errCodeReplace, lang)
}

// TranslationText
// @Description: 仅支持 - 分隔符
// @param l
// @param name
// @return string
func TranslationText(l string, text string) string {
	values := strings.Split(text, "-")
	if len(values) <= 0 {
		return text
	}
	newValue := MsgByValue(values[0], l)
	return strings.Replace(text, values[0], newValue, 1)
}
