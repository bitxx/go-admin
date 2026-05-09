// Package app
// @Description: 多语言消息管理，根据业务需要新增消息内容，标识内容勿操作
package lang

import (
	"errors"
	"fmt"
	"github.com/bitxx/logger/logbase"
	"strings"
)

// MsgByCode
// @Description: i18n
// @param errCode
// @param lang
// @return string
func MsgByCode(errCode int, lang string) string {
	switch lang {
	case "en":
		return EnLang.T(MsgInfo[errCode])
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
		return EnLang.T(value)
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
// @param errCodeReplace 最终需要给应用层返回的消息内容，这里传入消息码。若errCodeReplace=errCode或errCodeReplace<=0，则返回真实消息
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
