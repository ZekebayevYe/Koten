package subscriber

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/erazz/outage-service/pkg/email"
	natsgo "github.com/nats-io/nats.go"
)

type EmailSubscriber struct {
	nc     *natsgo.Conn
	sender *email.SMTPClient
}

func NewEmailSubscriber(nc *natsgo.Conn, sender *email.SMTPClient) *EmailSubscriber {
	return &EmailSubscriber{nc: nc, sender: sender}
}

func (s *EmailSubscriber) Subscribe(ctx context.Context) error {
	_, err := s.nc.Subscribe("notifications.created", func(msg *natsgo.Msg) {
		var notification struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Location    string `json:"location"`
		}

		if err := json.Unmarshal(msg.Data, &notification); err != nil {
			log.Printf("Error decoding notification: %v\n", err)
			return
		}

		to := os.Getenv("SMTP_USER")
		subject := "New Outage: " + notification.Title
		body := "Description: " + notification.Description + "\nLocation: " + notification.Location

		if err := s.sender.Send(to, subject, body); err != nil {
			log.Printf("Error sending email: %v\n", err)
		} else {
			log.Println("Notification email sent successfully")
		}
	})

	if err != nil {
		return err
	}

	return nil
}
