package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/erazz/outage-service/config"
	grpcsrv "github.com/erazz/outage-service/internal/delivery/grpc"
	"github.com/erazz/outage-service/internal/repository/mongo"
	"github.com/erazz/outage-service/internal/subscriber"
	"github.com/erazz/outage-service/internal/usecase"
	"github.com/erazz/outage-service/pkg/email"
	"github.com/erazz/outage-service/pkg/nats"
	natsgo "github.com/nats-io/nats.go"
)

func main() {
	_ = godotenv.Load()
	cfg := config.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	repo := mongo.New(ctx, cfg)
	pub := nats.New(cfg.NATSUrl)
	mail := email.New(cfg)
	svc := usecase.New(repo, pub, mail)
	s := grpc.NewServer()
	grpcsrv.Register(s, svc)

	go func() {
		nc, err := natsgo.Connect(os.Getenv("NATS_URL"))
		if err != nil {
			log.Fatal("failed to connect to NATS:", err)
		}

		smtpClient := email.NewSMTPClient()
		sub := subscriber.NewEmailSubscriber(nc, smtpClient)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := sub.Subscribe(ctx); err != nil {
			log.Fatal("failed to subscribe:", err)
		}

		select {}
	}()

	lis, err := net.Listen("tcp", ":"+cfg.ServicePort)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s.Serve(lis))
}
