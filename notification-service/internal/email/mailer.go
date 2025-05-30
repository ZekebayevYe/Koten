package email

import (
	"context"
	"github.com/ZekebayevYe/notification-service/internal/app"
	"github.com/mailersend/mailersend-go"
	"log"
	"os"
	"time"
)

type Mailer struct {
	client *mailersend.Mailersend
}

func NewMailer() *Mailer {
	apiKey := os.Getenv("MAILERSEND_API_KEY")
	return &Mailer{client: mailersend.NewMailersend(apiKey)}
}

func (m *Mailer) SendNotification(n app.Notification, to []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Mailer.SendNotification вызван:\n→ subject: %s\n→ message: %s\n→ to: %v\n", n.Title, n.Message, to)

	msg := m.client.Email.NewMessage()
	msg.SetFrom(mailersend.From{
		Name:  "Коммунсервис",
		Email: "no-reply@test-z0vklo6n7rxl7qrx.mlsender.net",
	})

	recipients := make([]mailersend.Recipient, 0, len(to))
	for _, e := range to {
		recipients = append(recipients, mailersend.Recipient{
			Email: e,
			Name:  "",
		})
	}
	msg.Recipients = recipients

	msg.Subject = n.Title
	msg.Text = n.Message
	msg.HTML = "<p>" + n.Message + "</p>"

	_, err := m.client.Email.Send(ctx, msg)
	if err != nil {
		log.Printf("MailerSend ошибка: %v", err)
		return
	}

	log.Println("MailerSend успешно отправил письмо")
}
