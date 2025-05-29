package nats

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type Producer struct {
	conn *nats.Conn
}

func NewProducer(url string) (*Producer, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Producer{conn: nc}, nil
}

func (p *Producer) Publish(ctx context.Context, subject string, data []byte) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.conn.Publish(subject, payload)
}

func (p *Producer) Close() {
	p.conn.Close()
}
