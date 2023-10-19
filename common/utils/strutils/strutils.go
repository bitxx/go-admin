package strutils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go-admin/common/global"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// HidePartStr
/**
 * @Description: 字符串中间替换为*
 * @param value
 * @param n
 * @return string
 */
func HidePartStr(value string, n int) string {
	if value == "" {
		return ""
	}
	startIndex := len(value)/2 - n/2
	replaceSymbol := "*"
	var builder strings.Builder
	for i, v := range value {
		if i >= startIndex-1 && i < startIndex+n {
			builder.WriteString(replaceSymbol)
		} else {
			builder.WriteString(string(v))
		}
	}
	return builder.String()
}

// IsNum 是否为整数
func IsNum(d decimal.Decimal) bool {
	if strings.Contains(d.String(), ".") {
		return false
	}
	return true
}

// GenerateValidateCode
// @Description: 随机生成6位数字验证码
// @return string
func GenerateValidateCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return rndCode
}

// VerifyEmailFormat
// @Description: 检测邮箱格式
// @param email
// @return bool
func VerifyEmailFormat(email string) bool {
	if email == "" {
		return false
	}
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	b := reg.MatchString(email)
	return b
}

// VersionOrdinal
// @Description:
// @param version
// @return string
// @return error
func VersionOrdinal(version string) (string, error) {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			return "", errors.New("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo), nil
}

type Address []byte

func IsMobile(mobile string) bool {
	//result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	//涉及到各国，因此，只要判断长度和是否纯数字
	if len(mobile) < 5 {
		return false
	}
	result, err := regexp.MatchString(`^[-+]?(([0-9]+)([.]([0-9]+))?|([.]([0-9]+))?)$`, mobile)
	if err != nil {
		return false
	}
	if !result {
		return false
	}
	return true
}

func Hmac(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func Base64ToImage(imageBase64 string) ([]byte, error) {
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)

	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}

	return filesRet, nil
}

func GetCurrentTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// slice去重
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func CompareHashAndPassword(e string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false
	}
	return true
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(global.TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(global.TrafficKey, requestId)
	}
	return requestId
}
