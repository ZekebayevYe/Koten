package domain

import "context"

type NotificationRepository interface {
	Create(ctx context.Context, n *Notification) (*Notification, error)
	GetByID(ctx context.Context, id string) (*Notification, error)
	List(ctx context.Context, f NotificationFilter) ([]Notification, error)
	Update(ctx context.Context, n *Notification) (*Notification, error)
	Delete(ctx context.Context, id string) error
}
