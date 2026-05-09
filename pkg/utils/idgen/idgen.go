package idgen

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

func UUID() string {
	return uuid.New().String()
}

func InviteId() string {
	return base64Captcha.RandText(6, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
