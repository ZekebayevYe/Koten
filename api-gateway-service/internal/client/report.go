package client

import (
	"api-gateway-service/config"
	reportPb "api-gateway-service/proto"
	"log"

	"google.golang.org/grpc"
)

func NewReportServiceClient(cfg *config.Config) reportPb.ReportServiceClient {
	conn, err := grpc.Dial(cfg.ReportServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to report-news-service: %v", err)
	}
	return reportPb.NewReportServiceClient(conn)
}
