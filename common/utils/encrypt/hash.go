package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

//
//  HashEncrypt
//  @Description: 将字符串单向加密
//  @param value
//  @return string
//  @return error
//
func HashEncrypt(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
