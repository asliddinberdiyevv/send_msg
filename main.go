package main

import (
	"fmt"
	"os"
	"send_msg/services"
)

func main() {

	// Sender data.
	fromEmail := os.Getenv("FROM_EMAIL")
	fromEmailpassword := os.Getenv("FROM_EMAIL_PASSWORD")
	toSendingUserList := []string{
		"javharbekernazarov47@gmail.com",
		"ruzikulovfurqat@gmail.com",
		"asliddinberdiyevv@gmail.com",
	}
	message := "This is a test email message."

	// Sending email.
	err := services.SenderGmail(fromEmail, fromEmailpassword, toSendingUserList, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
