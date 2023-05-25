package mail

import (
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {
	EMAIL_SENDER_NAME := "Bookhub"
	EMAIL_SENDER_ADDRESS := "asliddinberdiyevv@gmail.com"
	EMAIL_SENDER_PASSWORD := "dfwgtsubrcipuhsy"

	sender := NewGmailSender(EMAIL_SENDER_NAME, EMAIL_SENDER_ADDRESS, EMAIL_SENDER_PASSWORD)
	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	`
	to := []string{"ruzikulovfurqat@gmail.com"}
	attachFiles := []string{"./README.md"}

	err := sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	if err != nil {
		panic(err)
	}
}
