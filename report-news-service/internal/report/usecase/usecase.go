package usecase

import (
	"context"
	"time"

	"reportnewsservice/internal/report/repository"
	"reportnewsservice/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportUsecase interface {
	CreateReport(ctx context.Context, req *proto.CreateReportRequest) (*proto.Report, error)
	GetReportsByUser(ctx context.Context, userID string) ([]*proto.Report, error)
	EditReport(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error)
}

type reportUsecase struct {
	repo repository.ReportRepository
}

func NewReportUsecase(r repository.ReportRepository) ReportUsecase {
	return &reportUsecase{repo: r}
}

func (u *reportUsecase) CreateReport(ctx context.Context, req *proto.CreateReportRequest) (*proto.Report, error) {
	now := time.Now().Format(time.RFC3339)
	report := &proto.Report{
		ReportId:    primitive.NewObjectID().Hex(),
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := u.repo.Insert(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}

func (u *reportUsecase) GetReportsByUser(ctx context.Context, userID string) ([]*proto.Report, error) {
	return u.repo.FindByUser(ctx, userID)
}

func (u *reportUsecase) EditReport(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error) {
	return u.repo.Update(ctx, req)
}
