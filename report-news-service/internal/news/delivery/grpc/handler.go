package grpc

import (
	"context"

	newsusecase "reportnewsservice/internal/news/usecase"
	pb "reportnewsservice/proto"
)

type NewsHandler struct {
	pb.UnimplementedNewsServiceServer
	usecase newsusecase.NewsUsecase
}

func NewNewsHandler(uc newsusecase.NewsUsecase) *NewsHandler {
	return &NewsHandler{usecase: uc}
}

func (h *NewsHandler) GetAllNews(ctx context.Context, req *pb.Empty) (*pb.NewsList, error) {
	newsList, err := h.usecase.GetAllNews(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.NewsList{News: newsList}, nil
}
