package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
    user.Password = strings.TrimSpace(user.Password) // <- сначала trim

    hashedPassword, err := hash.HashPassword(user.Password) // <- потом хэш
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
	if err != nil {
		fmt.Println("[Login] user not found for email:", email, "error:", err)
		return "", errors.New("invalid email or password")
	}

	password = strings.TrimSpace(password) // trim входящий пароль

	fmt.Println("[Login] provided password:", password)
	fmt.Println("[Login] compare result:", hash.CheckPasswordHash(password, user.Password))

	fmt.Println("[Login] user found:", user.Email)
	fmt.Println("[Login] hashed password:", user.Password)

	if !hash.CheckPasswordHash(password, user.Password) {
		fmt.Println("[Login] password mismatch for email:", email)
		return "", errors.New("invalid email or password")
	}

	fmt.Println("[Login] password matched for email:", email)

	token, err := jwt.GenerateToken(user.Email, user.Role, u.config.JWTSecret, u.config.JWTExpiresIn)
	if err != nil {
		fmt.Println("[Login] token generation failed:", err)
		return "", err
	}

	fmt.Println("[Login] token generated for email:", email)
	return token, nil
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
