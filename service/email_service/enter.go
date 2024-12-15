package email_service

import (
	"blogv2/global"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"strings"
)

// 注册账号
func SendRegisterCode(to string, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】注册账号", em.SendNickName)
	text := fmt.Sprintf("你正在进行账号注册操作，这是你的验证码 %s 十分钟内有效", code)
	return SendEmail(to, subject, text)
}

// 重置密码
func SendResetPwdCode(to string, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】密码重置", em.SendNickName)
	text := fmt.Sprintf("你正在进行密码操作，这是你的验证码 %s 十分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendEmail(to, subject, text string) (err error) {
	em := global.Config.Email
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("%s <%s>", em.SendNickName, em.SendEmail)
	// 设置接收方的邮箱
	e.To = []string{to}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(text)
	//设置服务器相关的配置
	err = e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err != nil && !strings.Contains(err.Error(), "short response:") {
		logrus.Errorf(err.Error())
		return err
	}
	fmt.Println("发送成功")
	return nil
}
