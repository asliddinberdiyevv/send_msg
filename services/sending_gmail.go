package services

import (
	"net/smtp"
)

func SenderGmail(fromEmail string, fromEmailPassword string, toEmails []string, msg string) error {
	// Sender data.
	from := fromEmail
	password := fromEmailPassword

	// Receiver email address.
	to := toEmails

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
