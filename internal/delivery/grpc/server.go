package grpc

import (
	"google.golang.org/grpc"

	pb "github.com/erazz/outage-service/internal/delivery/grpc/proto"
	"github.com/erazz/outage-service/internal/domain"
)

func Register(s *grpc.Server, svc domain.NotificationService) {
	pb.RegisterNotificationServiceServer(s, NewHandler(svc))
}
