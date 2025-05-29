package client

import (
	"api-gateway-service/config"
	newsPb "api-gateway-service/proto"
	"log"

	"google.golang.org/grpc"
)

func NewNewsServiceClient(cfg *config.Config) newsPb.NewsServiceClient {
	conn, err := grpc.Dial(cfg.NewsServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to news service: %v", err)
	}
	return newsPb.NewNewsServiceClient(conn)
}
