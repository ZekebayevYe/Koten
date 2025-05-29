package broker

import (
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go" // Эта строка должна работать после установки пакета
)

type Publisher struct {
	nc *nats.Conn
}
type PublisherInterface interface {
	PublishDocumentEvent(event DocumentEvent) error
	Close()
}

// DocumentEvent представляет событие, связанное с документом
type DocumentEvent struct {
	EventType  string    `json:"event_type"` // "uploaded", "deleted", etc.
	DocumentID string    `json:"document_id"`
	UserID     string    `json:"user_id"`
	Filename   string    `json:"filename"`
	Timestamp  time.Time `json:"timestamp"`
}

// NewPublisher создает нового издателя для NATS
func NewPublisher(natsURL string) (*Publisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &Publisher{nc: nc}, nil
}

// PublishDocumentEvent публикует событие о документе
func (p *Publisher) PublishDocumentEvent(event DocumentEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	subject := "documents." + event.EventType
	return p.nc.Publish(subject, data)
}

// Close закрывает соединение с NATS
func (p *Publisher) Close() {
	if p.nc != nil {
		p.nc.Close()
	}
}

var _ PublisherInterface = (*Publisher)(nil)
