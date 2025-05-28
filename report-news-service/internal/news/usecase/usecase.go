package usecase

import (
	"context"
	"time"

	"reportnewsservice/internal/news/repository"
	"reportnewsservice/proto"
)

type NewsUsecase interface {
	GetAllNews(ctx context.Context) ([]*proto.News, error)
}

type newsUsecase struct {
	repo repository.NewsRepository
}

func NewNewsUsecase(r repository.NewsRepository) NewsUsecase {
	return &newsUsecase{repo: r}
}

func (u *newsUsecase) GetAllNews(ctx context.Context) ([]*proto.News, error) {
	return u.repo.FetchAll(ctx)
}

func (u *NewsUsecase) CreateNews(ctx context.Context, req *proto.CreateNewsRequest) (*proto.News, error) {
	news := &proto.News{
		NewsId:      req.NewsId,
		ImageUrl:    req.ImageUrl,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	err := u.repo.Insert(ctx, news)
	if err != nil {
		return nil, err
	}
	return news, nil
}
