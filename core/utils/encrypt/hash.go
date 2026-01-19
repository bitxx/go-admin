package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// HashEncrypt 将字符串单向加密
func HashEncrypt(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// MD5V 计算MD5哈希（兼容现有代码）
func MD5V(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256V 计算SHA256哈希
func SHA256V(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// SHA256VString 字符串版本的SHA256
func SHA256VString(s string) string {
	return SHA256V([]byte(s))
}

// SHA512V 计算SHA512哈希
func SHA512V(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// HashWithSalt 带盐值的哈希
func HashWithSalt(data []byte, salt string) string {
	h := sha256.New()
	h.Write(data)
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

// ShortHash 生成短哈希（取前n位）
func ShortHash(data []byte, length int) string {
	fullHash := SHA256V(data)
	if length <= 0 || length >= len(fullHash) {
		return fullHash
	}
	return fullHash[:length]
}
