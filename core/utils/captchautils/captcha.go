package captchautils

import (
	"go-admin/core/utils/idgen"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// SetStore 设置store
func SetStore(s base64Captcha.Store) {
	base64Captcha.DefaultMemStore = s
}

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverStringFunc() (id, b64s, answer string, err error) {
	e := configJsonBody{}
	e.Id = idgen.UUID()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{R: 240, G: 240, B: 246, A: 246}, nil, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return cap.Generate()
}

func DriverDigitFunc() (id, b64s, answer string, err error) {
	e := configJsonBody{}
	e.Id = idgen.UUID()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return cap.Generate()
}

// Verify 校验验证码
func Verify(id, code string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, code, clear)
}
