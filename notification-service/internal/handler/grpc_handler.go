package handler

import (
	"context"

	"github.com/ZekebayevYe/notification-service/internal/app"
	pb "github.com/ZekebayevYe/notification-service/internal/grpc"
	"github.com/ZekebayevYe/notification-service/internal/messaging"
	"github.com/ZekebayevYe/notification-service/internal/model"
)

type GRPCServer struct {
	svc *app.Service
	pb.UnimplementedNotificationServiceServer
}

func NewGRPCServer(svc *app.Service) *GRPCServer {
	return &GRPCServer{svc: svc}
}

func (g *GRPCServer) Subscribe(ctx context.Context, req *pb.EmailRequest) (*pb.Empty, error) {
	sub := model.Subscriber{
		Email:  req.Email,
		Street: req.Street,
		House:  req.House,
	}
	return &pb.Empty{}, g.svc.Subscribe(ctx, sub)
}

func (g *GRPCServer) Unsubscribe(ctx context.Context, req *pb.EmailRequest) (*pb.Empty, error) {
	return &pb.Empty{}, g.svc.Unsubscribe(ctx, req.Email)
}

func (g *GRPCServer) CreateNotification(ctx context.Context, req *pb.Notification) (*pb.Empty, error) {
	n := model.Notification{
		Title:   req.Title,
		Message: req.Message,
		SendAt:  req.SendAt,
		Street:  req.Street,
	}

	err := g.svc.CreateNotification(ctx, n)
	if err != nil {
		return nil, err
	}

	messaging.PublishNotification(n)

	return &pb.Empty{}, nil
}
func (g *GRPCServer) GetHistory(ctx context.Context, _ *pb.Empty) (*pb.NotificationList, error) {
	notifs, err := g.svc.GetAllNotifications(ctx)
	if err != nil {
		return nil, err
	}

	var result []*pb.Notification
	for _, n := range notifs {
		result = append(result, &pb.Notification{
			Title:   n.Title,
			Message: n.Message,
			SendAt:  n.SendAt,
			Street:  n.Street,
		})
	}

	return &pb.NotificationList{Items: result}, nil
}
