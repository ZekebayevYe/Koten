package usecase

import (
	"context"

	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"
)

type AuthUsecase struct {
	repo   domain.UserRepository
	config *config.Config
}

func NewAuthUsecase(repo domain.UserRepository, cfg *config.Config) *AuthUsecase {
	return &AuthUsecase{repo: repo, config: cfg}
}

func (u *AuthUsecase) Register(ctx context.Context, user *domain.User) (string, error) {
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	user.Role = "user"

	if err := u.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}

	return jwt.GenerateToken(user.Email, user.Role, u.config.JWTSecret, u.config.JWTExpiresIn)
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil || !hash.CheckPasswordHash(password, user.Password) {
		return "", err
	}

	return jwt.GenerateToken(user.Email, user.Role, u.config.JWTSecret, u.config.JWTExpiresIn)
}

func (u *AuthUsecase) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *AuthUsecase) UpdateProfile(ctx context.Context, email string, updated *domain.User) (*domain.User, error) {
	user, err := u.repo.UpdateUser(ctx, email, updated)
	if err != nil {
		return nil, err
	}
	return user, nil
}
