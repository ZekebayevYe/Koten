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
	fmt.Printf("[NewAuthServiceClient] Connecting to auth service at: %s\n", cfg.AuthServiceAddr)
	
	// 🔥 УЛУЧШЕНО: Добавляем опции подключения
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(), // Ждем подключения
		grpc.WithTimeout(10 * time.Second), // Таймаут подключения
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
	
	// Проверяем состояние подключения
	state := conn.GetState()
	fmt.Printf("[NewAuthServiceClient] Connection state: %s\n", state.String())
	
	if state != connectivity.Ready && state != connectivity.Idle {
		log.Printf("Warning: gRPC connection is not ready, state: %v", state)
	}
	
	client := pb.NewAuthServiceClient(conn)
	fmt.Println("[NewAuthServiceClient] Auth service client created successfully")
	
	return client
}

// 🔥 НОВОЕ: Функция для проверки подключения
func CheckConnection(client pb.AuthServiceClient) error {
	// Можно добавить health check если нужно
	return nil
}