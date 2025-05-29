package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZekebayevYe/documents-service/internal/broker"
	"github.com/ZekebayevYe/documents-service/internal/config"
	"github.com/ZekebayevYe/documents-service/internal/handler"
	"github.com/ZekebayevYe/documents-service/internal/storage"
	documentpb "github.com/ZekebayevYe/documents-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func connectMongo(ctx context.Context, uri string) (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetServerSelectionTimeout(30 * time.Second)
	clientOptions.SetConnectTimeout(30 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("MongoDB connection error: %w", err)
	}

	// Проверка подключения
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = client.Ping(ctxPing, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("MongoDB ping error: %w", err)
	}

	return client, client.Database("ums").Collection("documents"), nil
}

func main() {
	cfg := config.Load()

	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Подключение к MongoDB
	mongoClient, db, err := connectMongo(ctx, cfg.MongoURI)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("MongoDB disconnect error: %v", err)
		}
	}()

	// Создание publisher для NATS (может быть nil, если URL пустой)
	var pub broker.PublisherInterface
	if cfg.NatsURL != "" {
		publisher, err := broker.NewPublisher(cfg.NatsURL)
		if err != nil {
			log.Printf("NATS publisher creation failed: %v", err)
		} else {
			pub = publisher
			defer publisher.Close()
		}
	} else {
		log.Println("NATS_URL not set, publisher disabled")
	}

	grpcServer := grpc.NewServer()

	// Инициализация хранилища
	storageInstance := &storage.RealStorage{}

	documentpb.RegisterDocumentServiceServer(
		grpcServer,
		&handler.DocumentHandler{
			DB:        db,
			Publisher: pub,
			Storage:   storageInstance,
		},
	)

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	fmt.Printf("gRPC Document Service running on :%s\n", cfg.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
