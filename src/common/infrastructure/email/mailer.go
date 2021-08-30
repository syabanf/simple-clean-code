package email

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var (
	ConfigSmtpHost     = os.Getenv("CONFIG_SMTP_HOST")
	ConfigSmtpPort     = 587
	ConfigSenderName   = os.Getenv("CONFIG_SENDER_NAME")
	ConfigAuthEmail    = os.Getenv("CONFIG_AUTH_EMAIL")
	ConfigAuthPassword = os.Getenv("CONFIG_AUTH_PASSWORD")
)

func SendMail(subject string, recipient []string, body string) (err error) {
	godotenv.Load()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", ConfigSenderName)
	mailer.SetHeader("To", recipient...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		ConfigSmtpHost,
		ConfigSmtpPort,
		ConfigSenderName,
		ConfigAuthEmail,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		err = fmt.Errorf("Failed send mail")
		log.Println("Failed send mail")
		return
	}

	return
}
