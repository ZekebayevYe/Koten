package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"
)

type AuthUsecase struct {
	repo      domain.UserRepository
	publisher domain.EventPublisher
	config    *config.Config
}

func NewAuthUsecase(repo domain.UserRepository, publisher domain.EventPublisher, cfg *config.Config) *AuthUsecase {
	return &AuthUsecase{repo: repo, publisher: publisher, config: cfg}
}

func (u *AuthUsecase) Register(ctx context.Context, user *domain.User) (string, error) {
	user.Password = strings.TrimSpace(user.Password)
	fmt.Println("[Register] password before hashing:", user.Password)
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		fmt.Println("[Register] hash error:", err)
		return "", err
	}
	fmt.Println("[Register] hashed password:", hashedPassword)
	user.Password = hashedPassword
	user.Role = "user"

	if err := u.repo.CreateUser(ctx, user); err != nil {
		fmt.Println("[Register] create user error:", err)
		return "", err
	}

	event := map[string]string{
		"email": user.Email,
		"role":  user.Role,
	}
	eventData, _ := json.Marshal(event)

	if err := u.publisher.Publish(ctx, "user.registered", eventData); err != nil {
		fmt.Println("[Register] failed to publish user.registered:", err)
	}

	token, err := jwt.GenerateToken(user.Email, user.Role)
	if err != nil {
		fmt.Println("[Register] token generation error:", err)
		return "", err
	}
	fmt.Println("[Register] token generated:", token)
	return token, nil
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("[Login] user not found for email:", email, "error:", err)
		return "", errors.New("invalid email or password")
	}

	fmt.Println("[Login] raw password (before trim):", password)
	password = strings.TrimSpace(password)
	fmt.Println("[Login] trimmed password:", password)
	fmt.Println("[Login] stored hash:", user.Password)

	if !hash.CheckPasswordHash(password, user.Password) {
		fmt.Println("[Login] password mismatch for email:", email)
		return "", errors.New("invalid email or password")
	}

	fmt.Println("[Login] password matched for email:", email)

	token, err := jwt.GenerateToken(user.Email, user.Role)
	if err != nil {
		fmt.Println("[Login] token generation failed:", err)
		return "", err
	}

	fmt.Println("[Login] token generated:", token)
	return token, nil
}

func (u *AuthUsecase) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("[GetProfile] error:", err)
		return nil, err
	}
	return user, nil
}

func (u *AuthUsecase) UpdateProfile(ctx context.Context, email string, updated *domain.User) (*domain.User, error) {
	user, err := u.repo.UpdateUser(ctx, email, updated)
	if err != nil {
		fmt.Println("[UpdateProfile] error:", err)
		return nil, err
	}
	return user, nil
}
