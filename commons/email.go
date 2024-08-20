package common

import (
	"fmt"
	"net/smtp"
	"log"
	config "ginTailwindcssBase/config"
)


//发送验证码，对对方邮箱，验证码
func Sendmail(dfemail string,code string) {
	title := "请接收您的验证码"
	body := fmt.Sprintf(`<div style="background-color: #f2f2f2; padding: 20px;">
	<h2 style="color: #333; font-size: 24px; margin-bottom: 20px;">以下是您的验证码：</h2>
	<div style="background-color: #fff; border: 1px solid green; color:green; border-radius: 5px; padding: 10px; font-size: 24px; font-weight: bold; text-align: center; margin-bottom: 20px;">
	%s
	</div>
	<p style="color: #666; font-size: 16px;">此电子邮件地址仅用于发送通知，无法接收电子邮件，请勿回复此邮件。</p>
	</div>`, code)
	tuisongemail(dfemail,title,body)
}

// 参数：对方邮箱，标题，内容
func tuisongemail(dfemail string,title string, body string) {
	auth := smtp.PlainAuth("", config.Config.EmailUser, config.Config.EmailPassword, "smtp.qq.com")
	// 邮件内容
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "From: "+config.Config.EmailUser+"@qq.com\n" + "Subject: " + title + "\n" + mime + "\n" + body

	// 发送邮件
	err := smtp.SendMail("smtp.qq.com:587", auth, config.Config.EmailUser+"@qq.com", []string{dfemail}, []byte(msg))
	if err != nil {
		log.Print("QQ邮箱发送失败:", err)
		return
	}
}