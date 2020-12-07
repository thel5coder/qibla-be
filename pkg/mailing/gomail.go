package mailing

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type GoMailConfig struct {
	SMTPHost string
	SMTPPort int
	Sender   string
	Password string
}

func(goMailConfig GoMailConfig) SendGoMail(to,subject,body string){
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", goMailConfig.Sender)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		goMailConfig.SMTPHost,
		goMailConfig.SMTPPort,
		goMailConfig.Sender,
		goMailConfig.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Print(err.Error())
	}
}
