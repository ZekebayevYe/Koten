package domain

import "context"

type NotificationService interface {
	Create(ctx context.Context, n Notification) (Notification, error)
	Get(ctx context.Context, id string) (Notification, error)
	List(ctx context.Context, f NotificationFilter) ([]Notification, error)
	Update(ctx context.Context, n Notification) (Notification, error)
	Delete(ctx context.Context, id string) error
}
