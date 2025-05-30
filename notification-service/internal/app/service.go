package app

import (
	"context"
	"errors"
	"github.com/ZekebayevYe/notification-service/internal/model"
	"time"
)

var (
	ErrAlreadySubscribed = errors.New("уже подписан")
)

type Notification = model.Notification

type Service struct {
	repo     Repository
	cacheTTL time.Duration
	mailer   Mailer
}

type Repository interface {
	AddSubscriber(context.Context, model.Subscriber) error
	RemoveSubscriber(context.Context, string) error
	ListSubscribers(context.Context) ([]string, error)
	ListSubscribersByStreet(context.Context, string) ([]string, error)
	SaveNotification(context.Context, Notification) error
	GetAllNotifications(context.Context) ([]model.Notification, error)
}

type Mailer interface {
	SendNotification(Notification, []string)
}

func (s *Service) GetAllNotifications(ctx context.Context) ([]model.Notification, error) {
	return s.repo.GetAllNotifications(ctx)
}

func NewService(r Repository, m Mailer, cacheTTL time.Duration) *Service {
	return &Service{repo: r, mailer: m, cacheTTL: cacheTTL}
}

func (s *Service) Subscribe(ctx context.Context, sub model.Subscriber) error {
	return s.repo.AddSubscriber(ctx, sub)
}

func (s *Service) Unsubscribe(ctx context.Context, email string) error {
	return s.repo.RemoveSubscriber(ctx, email)
}

func (s *Service) CreateNotification(ctx context.Context, n Notification) error {
	if err := s.repo.SaveNotification(ctx, n); err != nil {
		return err
	}
	return nil
}
