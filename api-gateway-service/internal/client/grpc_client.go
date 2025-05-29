package client

import (
	"api-gateway-service/config"
	pb "api-gateway-service/proto"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/keepalive"
)

func NewAuthServiceClient(cfg *config.Config) pb.AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(), 
		grpc.WithTimeout(10 * time.Second), 
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}),
	}
	
	conn, err := grpc.Dial(cfg.AuthServiceAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to auth-service at %s: %v", cfg.AuthServiceAddr, err)
	}
	
	state := conn.GetState()
	fmt.Printf("[NewAuthServiceClient] Connection state: %s\n", state.String())
	
	if state != connectivity.Ready && state != connectivity.Idle {
		log.Printf("Warning: gRPC connection is not ready, state: %v", state)
	}
	
	client := pb.NewAuthServiceClient(conn)
	fmt.Println("[NewAuthServiceClient] Auth service client created successfully")
	
	return client
}

func CheckConnection(client pb.AuthServiceClient) error {
	return nil
}