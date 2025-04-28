package tests

import (
	"context"
	"testing"

	pb "github.com/erazz/outage-service/internal/delivery/grpc/proto"
	"google.golang.org/grpc"
)

func TestHealth(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewNotificationServiceClient(conn)
	_, err = c.HealthCheck(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
}
