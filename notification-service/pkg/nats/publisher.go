package nats

import (
	"context"

	natsgo "github.com/nats-io/nats.go"
)

type Publisher interface {
	Publish(ctx context.Context, subject string, data []byte) error
}

type publisher struct {
	nc *natsgo.Conn
}

func New(url string) Publisher {
	nc, _ := natsgo.Connect(url)
	return &publisher{nc: nc}
}

func (p *publisher) Publish(_ context.Context, subject string, data []byte) error {
	return p.nc.Publish(subject, data)
}
