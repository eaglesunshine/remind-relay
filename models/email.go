package models

import (
	"gopkg.in/gomail.v2"
)

//EmailConfig 发生邮件
type EmailConfig struct {
	MailHost string // 邮件服务器地址
	MailPort int    // 端口
	MailUser string // 发送邮件用户账号
	MailPwd  string // 授权密码
}

/*
   title 使用gomail发送邮件
   @param []string mailAddress 收件人邮箱
   @param string subject 邮件主题
   @param string body 邮件内容
   @return error
*/

//SendGoMail 发送邮件
func SendGoMail(mailAddress []string, subject string, body string) error {
	//获取远程配置
	var conf EmailConfig
	//if err := golib.LoadConfigStruct("EmailConfig", &conf, true); err != nil {
	//	return err
	//}

	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "gomail"
	m.SetHeader("From", nickname+"<"+conf.MailUser+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(conf.MailHost, conf.MailPort, conf.MailUser, conf.MailPwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}

//SendEmail 发送邮件入口
func SendEmail(to, title, content string) error {
	return SendGoMail([]string{to}, title, content)
}
