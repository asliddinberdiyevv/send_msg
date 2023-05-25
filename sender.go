package mail

type EmailSender struct {
	SendEmail (
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
} 

