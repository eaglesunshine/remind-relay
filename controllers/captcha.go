package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"remind-relay/models"
)

// CaptchaController 生成校验码数据
type CaptchaController struct {
	beego.Controller
}

// Get 接收get请求
func (this *CaptchaController) Get() {
	//数字验证码配置
	var configD = base64Captcha.DriverDigit{
		Height:   40,
		Width:    130,
		MaxSkew:  0,
		DotCount: 0,
		Length:   5,
	}

	////声音验证码配置
	//var configA = base64Captcha.DriverAudio{
	//	Length:   6,
	//	Language: "zh",
	//}
	//
	////字符验证码配置
	//var configS = base64Captcha.DriverString{
	//	Height: 60,
	//	Width:  240,
	//}
	//
	////汉字验证码配置
	//var configC = base64Captcha.DriverChinese{}

	//数学验证码配置
	rgbaColor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	var configM = base64Captcha.DriverMath{
		Height:          40,
		Width:           140,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor:         &rgbaColor,
		Fonts:           fonts,
	}

	param := models.ConfigJsonBody{
		CaptchaType: "math",
		DriverDigit: &configD,
		DriverMath:  &configM,
	}

	ret, err := models.GenerateCaptchaHandler(param)
	if err != nil {
		SetErrorReturn(this, EC_INNER_ERR, err)
		return
	}

	this.Data["json"] = map[string]interface{}{
		"status": 0,
		"data":   ret,
		"msg":    "",
	}
	this.ServeJSON()
}
