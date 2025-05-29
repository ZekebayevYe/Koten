package main

import (
	"context"
	"log"
	"net"
	"strings"

	"auth-service/config"
	delivery "auth-service/internal/delivery"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"
	"auth-service/pkg/jwtutil"
	"auth-service/pkg/nats"
	pb "auth-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 🔥 ДОБАВЛЕНО: JWT interceptor для проверки токенов
func jwtInterceptor(jwtSecret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Пропускаем аутентификацию для login и register
		if strings.Contains(info.FullMethod, "LoginUser") || strings.Contains(info.FullMethod, "RegisterUser") {
			return handler(ctx, req)
		}

		// Извлекаем metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token not found")
		}

		// Получаем токен
		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token not found")
		}

		authHeader := authHeaders[0]
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header format")
		}

		tokenStr := parts[1]

		// Парсим JWT токен
		email, role, err := jwtutil.ParseToken(tokenStr, jwtSecret)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		// Добавляем email и role в контекст
		ctx = context.WithValue(ctx, "email", email)
		ctx = context.WithValue(ctx, "role", role)

		return handler(ctx, req)
	}
}

func main() {
	cfg := config.Load()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	natsProducer, err := nats.NewProducer(cfg.NATSUrl)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer natsProducer.Close()

	repo := repository.NewMongoUserRepository(client.Database(cfg.MongoDB), "users")
	usecase := usecase.NewAuthUsecase(repo, natsProducer, cfg)
	handler := &delivery.Handler{Usecase: usecase, Cfg: cfg}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen on :50051:", err)
	}

	// 🔥 ИЗМЕНЕНО: добавлен JWT interceptor
	s := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor(cfg.JWTSecret)),
	)
	pb.RegisterAuthServiceServer(s, handler)

	log.Println("AuthService started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC:", err)
	}
}
