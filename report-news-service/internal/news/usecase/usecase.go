package usecase

import (
	"context"

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
