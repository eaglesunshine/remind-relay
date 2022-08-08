package models

import (
	"github.com/mojocn/base64Captcha"
)

//ConfigJsonBody 请求参数
type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptchaHandler 生成验证码数据
func GenerateCaptchaHandler(param ConfigJsonBody) (map[string]interface{}, error) {
	//parse request parameters
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()

	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
		return body, err
	}

	return body, nil
}

// CaptchaVerifyHandle 校验
func CaptchaVerifyHandle(id, value string) bool {
	return store.Verify(id, value, true)
}
