package email

import (
	"net/smtp"

	"github.com/erazz/outage-service/config"
)

type Sender interface {
	Send(to []string, subject, body string) error
}

type sender struct {
	auth smtp.Auth
	host string
	from string
}

func New(cfg config.Config) Sender {
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	return &sender{auth: auth, host: cfg.SMTPHost + ":" + cfg.SMTPPort, from: cfg.SMTPFrom}
}

func (s *sender) Send(to []string, subject, body string) error {
	msg := "From: " + s.from + "\r\n" +
		"To: " + s.from + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" + body
	return smtp.SendMail(s.host, s.auth, s.from, to, []byte(msg))
}
