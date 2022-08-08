package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"os"
	"regexp"
	"remind-relay/models"
)

// ValidateCodeController 注册生成校验码数据(通过邮箱发送)
type ValidateCodeController struct {
	beego.Controller
}

//验证码邮件模版
var message = `
    <p> 尊敬的itango用户：</p>
		<p style="text-indent:2em">验证码：${code}，请在5分钟内完成验证。</p>
	`

// Post 接收Post请求
func (this *ValidateCodeController) Post() {
	//1、验证码机制，拦截恶意调用
	id := this.Input().Get("CaptchaId")
	value := this.Input().Get("CaptchaValue")

	if len(id) == 0 || len(value) == 0 || models.CaptchaVerifyHandle(id, value) == false {
		SetErrorReturn(this, EC_INNER_ERR, errors.New("图形验证码错误！"))
		return
	}

	//获取邮箱号
	username := this.Input().Get("Username")
	if len(username) == 0 {
		SetErrorReturn(this, EC_INNER_ERR, errors.New("邮箱号为空！"))
		return
	}
	if VerifyEmailFormat(username) == false {
		SetErrorReturn(this, EC_INNER_ERR, errors.New("邮箱号格式错误！"))
		return
	}

	//生成6位随机验证码
	code, err := models.GenValidateCode(6, username)
	if err != nil {
		SetErrorReturn(this, EC_INNER_ERR, err)
		return
	}

	//构造邮件body
	mp := map[string]string{"code": code}
	body := os.Expand(message, func(k string) string { return mp[k] })

	//向邮箱号发送邮件
	err = models.SendEmail(username, "验证码", body)
	if err != nil {
		SetErrorReturn(this, EC_INNER_ERR, err)
		return
	}

	this.Data["json"] = map[string]interface{}{
		"status": 0,
		"data":   "验证码已发送至您的邮箱，请注意查收",
		"msg":    "",
	}
	this.ServeJSON()
}

//VerifyEmailFormat 校验邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
