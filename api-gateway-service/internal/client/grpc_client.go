package client

import (
	"api-gateway-service/config"
	pb "api-gateway-service/proto"
	"log"

	"google.golang.org/grpc"
)

func NewAuthServiceClient(cfg *config.Config) pb.AuthServiceClient {
	conn, err := grpc.Dial(cfg.AuthServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}
	return pb.NewAuthServiceClient(conn)
}
