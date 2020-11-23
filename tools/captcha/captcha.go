/*
@Time : 2020/11/4 下午10:53
@Author : hoastar
@File : captcha
@Software: GoLand
*/

package captcha

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// configJsonBody json request body
type configJsonBody struct {
	Id	string
	CaptchaType string
	VerifyValue string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// DriverStringFunc generate stringCaptcha
func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	// random string: cat /dev/urandom | tr -dc a-z0-9 | head -c 35
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "90abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

// DriverStringFunc generate digitCaptcha
func DriverDigitFunc() (id, base64 string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}






















