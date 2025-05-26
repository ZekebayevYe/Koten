package grpc

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/erazz/outage-service/internal/delivery/grpc/proto"
	"github.com/erazz/outage-service/internal/domain"
)

type handler struct {
	pb.UnimplementedNotificationServiceServer
	svc domain.NotificationService
	sub chan *pb.Notification
}

func NewHandler(svc domain.NotificationService) *handler {
	return &handler{svc: svc, sub: make(chan *pb.Notification, 100)}
}

func toProto(n domain.Notification) *pb.Notification {
	return &pb.Notification{
		Id:           n.ID.Hex(),
		Title:        n.Title,
		Description:  n.Description,
		ResourceType: n.ResourceType,
		Location:     n.Location,
		StartTime:    n.StartTime.Unix(),
		EndTime:      n.EndTime.Unix(),
		CreatedAt:    n.CreatedAt.Unix(),
	}
}

func (h *handler) CreateNotification(ctx context.Context, r *pb.CreateNotificationRequest) (*pb.Notification, error) {
	n, err := h.svc.Create(ctx, domain.Notification{
		Title:        r.Title,
		Description:  r.Description,
		ResourceType: r.ResourceType,
		Location:     r.Location,
		StartTime:    time.Unix(r.StartTime, 0),
		EndTime:      time.Unix(r.EndTime, 0),
	})
	if err != nil {
		return nil, err
	}
	p := toProto(n)
	select {
	case h.sub <- p:
	default:
	}
	return p, nil
}

func (h *handler) GetNotification(ctx context.Context, r *pb.NotificationId) (*pb.Notification, error) {
	n, err := h.svc.Get(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return toProto(n), nil
}

func (h *handler) ListNotifications(ctx context.Context, r *pb.ListNotificationsRequest) (*pb.Notifications, error) {
	from := time.Unix(r.DateFrom, 0)
	to := time.Unix(r.DateTo, 0)
	list, err := h.svc.List(ctx, domain.NotificationFilter{
		ResourceType: r.ResourceType,
		Location:     r.Location,
		DateFrom:     from,
		DateTo:       to,
	})
	if err != nil {
		return nil, err
	}
	res := &pb.Notifications{}
	for _, n := range list {
		res.Items = append(res.Items, toProto(n))
	}
	return res, nil
}

func (h *handler) UpdateNotification(ctx context.Context, r *pb.UpdateNotificationRequest) (*pb.Notification, error) {
	oid, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, err
	}
	n, err := h.svc.Update(ctx, domain.Notification{
		ID:           oid,
		Title:        r.Title,
		Description:  r.Description,
		ResourceType: r.ResourceType,
		Location:     r.Location,
		StartTime:    time.Unix(r.StartTime, 0),
		EndTime:      time.Unix(r.EndTime, 0),
	})
	if err != nil {
		return nil, err
	}
	return toProto(n), nil
}

func (h *handler) DeleteNotification(ctx context.Context, r *pb.NotificationId) (*pb.Empty, error) {
	err := h.svc.Delete(ctx, r.Id)
	return &pb.Empty{}, err
}

func (h *handler) GetNotificationsByDate(ctx context.Context, r *pb.ListNotificationsRequest) (*pb.Notifications, error) {
	return h.ListNotifications(ctx, r)
}

func (h *handler) GetNotificationsByLocation(ctx context.Context, r *pb.ListNotificationsRequest) (*pb.Notifications, error) {
	return h.ListNotifications(ctx, r)
}

func (h *handler) GetNotificationsByResource(ctx context.Context, r *pb.ListNotificationsRequest) (*pb.Notifications, error) {
	return h.ListNotifications(ctx, r)
}

func (h *handler) HealthCheck(ctx context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (h *handler) Ping(ctx context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (h *handler) SubscribeNotifications(_ *pb.Empty, stream pb.NotificationService_SubscribeNotificationsServer) error {
	for n := range h.sub {
		if err := stream.Send(n); err != nil {
			return err
		}
	}
	return nil
}

func (h *handler) AckNotification(ctx context.Context, _ *pb.NotificationId) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
