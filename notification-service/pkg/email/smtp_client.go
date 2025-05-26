package email

import (
	"net/smtp"
	"os"
)

type SMTPClient struct {
	auth smtp.Auth
	from string
	addr string
}

func NewSMTPClient() *SMTPClient {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", user, pass, host)
	addr := host + ":" + port

	return &SMTPClient{
		auth: auth,
		from: user,
		addr: addr,
	}
}

func (c *SMTPClient) Send(to, subject, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	return smtp.SendMail(c.addr, c.auth, c.from, []string{to}, msg)
}
