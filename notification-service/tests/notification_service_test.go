package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/erazz/outage-service/internal/domain"
	"github.com/erazz/outage-service/internal/usecase"
)

type repoStub struct{ n domain.Notification }

func (r *repoStub) Create(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	r.n = *n
	return n, nil
}
func (r *repoStub) GetByID(ctx context.Context, id string) (*domain.Notification, error) {
	return &r.n, nil
}
func (r *repoStub) List(ctx context.Context, f domain.NotificationFilter) ([]domain.Notification, error) {
	return []domain.Notification{r.n}, nil
}
func (r *repoStub) Update(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	r.n = *n
	return n, nil
}
func (r *repoStub) Delete(ctx context.Context, id string) error { return nil }

type pubStub struct{}

func (p *pubStub) Publish(ctx context.Context, s string, b []byte) error { return nil }

type mailStub struct{}

func (m *mailStub) Send(to []string, s, b string) error { return nil }

func TestCreate(t *testing.T) {
	svc := usecase.New(&repoStub{}, &pubStub{}, &mailStub{})
	n, err := svc.Create(context.TODO(), domain.Notification{Title: "t"})
	assert.NoError(t, err)
	assert.Equal(t, "t", n.Title)
}
