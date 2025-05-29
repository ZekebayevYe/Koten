package main

import (
	"context"
	"log"
	"net"

	"auth-service/config"
	delivery "auth-service/internal/delivery"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"
	pb "auth-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	repo := repository.NewMongoUserRepository(client.Database(cfg.MongoDB), "users")
	usecase := usecase.NewAuthUsecase(repo, cfg)
	handler := &delivery.Handler{Usecase: usecase, Cfg: cfg}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, handler)

	log.Println("AuthService started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
