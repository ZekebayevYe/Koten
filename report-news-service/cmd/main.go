package main

import (
	"context"
	"log"
	"net"
	"time"

	newsGrpc "reportnewsservice/internal/news/delivery/grpc"
	newsRepo "reportnewsservice/internal/news/repository"
	newsUsecase "reportnewsservice/internal/news/usecase"
	reportGrpc "reportnewsservice/internal/report/delivery/grpc"
	reportRepo "reportnewsservice/internal/report/repository"
	reportUsecase "reportnewsservice/internal/report/usecase"
	"reportnewsservice/proto"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// Logger setup
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// MongoDB setup
	mongoURI := "mongodb+srv://erazzzul:HEu9kOvHAQmVpOxa@yera.gfdef.mongodb.net/?retryWrites=true&w=majority&appName=yera"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("reportnews")

	// NATS setup
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Prometheus
	go func() {
		httpRouter := gin.Default()
		httpRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))
		httpRouter.Run(":9090")
	}()

	// Report service wiring
	rRepo := reportRepo.NewReportRepository(db)
	rUsecase := reportUsecase.NewReportUsecase(rRepo)
	rHandler := reportGrpc.NewReportHandler(rUsecase)

	// News service wiring
	nRepo := newsRepo.NewNewsRepository(db)
	nUsecase := newsUsecase.NewNewsUsecase(nRepo)
	nHandler := newsGrpc.NewNewsHandler(nUsecase)

	// gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		proto.RegisterReportServiceServer(grpcServer, rHandler)
		proto.RegisterNewsServiceServer(grpcServer, nHandler)
		log.Println("gRPC server running at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// HTTP server stub
	r := gin.Default()
	// You can add HTTP endpoints here if needed
	r.Run(":8080")
}
