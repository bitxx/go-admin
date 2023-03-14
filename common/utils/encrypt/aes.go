package encrypt

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
)

var defaultKey = []byte{159, 88, 162, 9, 70, 180, 126, 25, 0, 3, 236, 113, 108, 28, 69, 125}

func AesEncryptDefault(text string) (string, error) {
	if text == "" {
		return "", nil
	}
	return AesEncrypt(text, defaultKey)
}

func AesDecryptDefault(text string) (string, error) {
	if text == "" {
		return "", nil
	}
	return AesDecrypt(text, defaultKey)
}

func AesEncrypt(v string, k []byte) (string, error) {
	value := []byte(v)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	blocksize := block.BlockSize()
	valueBytes := value

	fillsize := blocksize - len(valueBytes)%blocksize
	repeat := bytes.Repeat([]byte{byte(fillsize)}, fillsize)
	valueBytes = append(valueBytes, repeat...)

	result := make([]byte, len(valueBytes))

	temp := result
	for len(valueBytes) > 0 {
		block.Encrypt(temp, valueBytes[:blocksize])
		valueBytes = valueBytes[blocksize:]
		temp = temp[blocksize:]
	}
	return hex.EncodeToString(result), nil
}

func AesDecrypt(v string, k []byte) (string, error) {
	value, _ := hex.DecodeString(v)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	size := len(value)
	result := make([]byte, size)
	blocksize := block.BlockSize()
	if size%blocksize != 0 {
		return "", errors.New("待解密数据异常")
	}
	temp := result
	for len(value) > 0 {
		block.Decrypt(temp, value[:blocksize])
		value = value[blocksize:]
		temp = temp[blocksize:]
	}
	count := 0
	for i := size - 1; i >= 0; i-- {
		if result[i] > 16 { //尾部空格
			break
		}
		count++
	}
	sub := size - count
	if sub < 0 {
		sub = 0
	}
	result = result[:sub]
	return string(result), nil
}
