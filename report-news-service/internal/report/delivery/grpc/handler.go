package grpc

import (
	"context"

	"reportnewsservice/internal/report/usecase"
	"reportnewsservice/proto"
)

type ReportHandler struct {
	proto.UnimplementedReportServiceServer
	usecase usecase.ReportUsecase
}

func NewReportHandler(u usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{usecase: u}
}

func (h *ReportHandler) CreateReport(ctx context.Context, req *proto.CreateReportRequest) (*proto.ReportResponse, error) {
	report, err := h.usecase.CreateReport(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.ReportResponse{Report: report}, nil
}

func (h *ReportHandler) GetReportsByUser(ctx context.Context, req *proto.UserRequest) (*proto.ReportsResponse, error) {
	reports, err := h.usecase.GetReportsByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.ReportsResponse{Reports: reports}, nil
}

func (h *ReportHandler) EditReport(ctx context.Context, req *proto.EditReportRequest) (*proto.ReportResponse, error) {
	report, err := h.usecase.EditReport(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.ReportResponse{Report: report}, nil
}
