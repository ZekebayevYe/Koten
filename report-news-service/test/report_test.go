package usecase_test

import (
	"context"
	"testing"

	"reportnewsservice/internal/report/usecase"
	"reportnewsservice/proto"

	"github.com/stretchr/testify/assert"
)

type mockReportRepo struct {
	inserted *proto.Report
}

func (m *mockReportRepo) Insert(ctx context.Context, r *proto.Report) error {
	m.inserted = r
	return nil
}

func (m *mockReportRepo) FindByUser(ctx context.Context, userID string) ([]*proto.Report, error) {
	return []*proto.Report{}, nil
}

func (m *mockReportRepo) Update(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error) {
	return &proto.Report{
		ReportId:    req.ReportId,
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		UpdatedAt:   "now",
	}, nil
}

func TestCreateReport(t *testing.T) {
	mockRepo := &mockReportRepo{}
	usecase := usecase.NewReportUsecase(mockRepo)

	req := &proto.CreateReportRequest{
		UserId:      "user123",
		Title:       "Test Title",
		Description: "Test Description",
	}

	created, err := usecase.CreateReport(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Title, created.Title)
	assert.Equal(t, req.Description, created.Description)
	assert.Equal(t, req.UserId, created.UserId)
	assert.Equal(t, mockRepo.inserted, created)
}
