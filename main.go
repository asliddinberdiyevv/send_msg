package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {

	// Sender data.
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("FROM_EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		"javharbekernazarov47@gmail.com",
		"ruzikulovfurqat@gmail.com",
		"asliddinberdiyevv@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
