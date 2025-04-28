package usecase

import (
	"context"
	"encoding/json"

	"github.com/erazz/outage-service/internal/domain"
	"github.com/erazz/outage-service/pkg/email"
	"github.com/erazz/outage-service/pkg/nats"
)

type service struct {
	repo domain.NotificationRepository
	pub  nats.Publisher
	mail email.Sender
}

func New(r domain.NotificationRepository, p nats.Publisher, m email.Sender) domain.NotificationService {
	return &service{repo: r, pub: p, mail: m}
}

func (s *service) Create(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	res, err := s.repo.Create(ctx, &n)
	if err != nil {
		return domain.Notification{}, err
	}
	b, _ := json.Marshal(res)
	_ = s.pub.Publish(ctx, "outages.created", b)
	_ = s.mail.Send([]string{}, "new outage", string(b))
	return *res, nil
}

func (s *service) Get(ctx context.Context, id string) (domain.Notification, error) {
	n, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Notification{}, err
	}
	return *n, nil
}

func (s *service) List(ctx context.Context, f domain.NotificationFilter) ([]domain.Notification, error) {
	return s.repo.List(ctx, f)
}

func (s *service) Update(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	res, err := s.repo.Update(ctx, &n)
	if err != nil {
		return domain.Notification{}, err
	}
	return *res, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
