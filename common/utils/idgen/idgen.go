package idgen

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"strconv"
)

func Id() string {
	//返回128位的高位部分
	return strconv.FormatUint(idgenerator.NewRandom128().TraceID().High, 10)
}

func UUID() string {
	return uuid.New().String()
}

func InviteId() string {
	return base64Captcha.RandText(6, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
