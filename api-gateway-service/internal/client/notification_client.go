package client

import (
	"api-gateway-service/config"
	notificationpb "api-gateway-service/proto/notification"
	"log"

	"google.golang.org/grpc"
)

func NewNotificationServiceClient(cfg *config.Config) notificationpb.NotificationServiceClient {
	conn, err := grpc.Dial(cfg.NotificationServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to notification-service: %v", err)
	}
	return notificationpb.NewNotificationServiceClient(conn)
}
