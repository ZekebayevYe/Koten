package usecase

import (
	"auth-service/config"
	"auth-service/internal/domain"
	ucCache "auth-service/internal/usecase/cache"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type AuthUsecase struct {
	repo      domain.UserRepository
	publisher domain.EventPublisher
	config    *config.Config
	cache     *ucCache.UserCache
}

func NewAuthUsecase(repo domain.UserRepository, publisher domain.EventPublisher, cfg *config.Config, cache *ucCache.UserCache) *AuthUsecase {
	return &AuthUsecase{
		repo:      repo,
		publisher: publisher,
		config:    cfg,
		cache:     cache,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, user *domain.User) (string, error) {
	user.Password = strings.TrimSpace(user.Password)

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()


	if err := u.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}

	event := map[string]interface{}{
		"email":      user.Email,
		"role":       user.Role,
		"full_name":  user.FullName,
		"house":      user.House,
		"street":     user.Street,
		"apartment":  user.Apartment,
		"created_at": user.CreatedAt,
	}

	eventData, err := json.Marshal(event)
	if err != nil {
	} else {
		if err := u.publisher.Publish(ctx, "user.registered", eventData); err != nil {
		} else {
		}
	}

	token, err := jwt.GenerateToken(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	password = strings.TrimSpace(password)
	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	loginEvent := map[string]interface{}{
		"email":     user.Email,
		"role":      user.Role,
		"action":    "login",
		"timestamp": fmt.Sprintf("%d", ctx.Value("timestamp")),
	}

	if eventData, err := json.Marshal(loginEvent); err == nil {
		if err := u.publisher.Publish(ctx, "user.login", eventData); err != nil {
		} else {
		}
	}

	token, err := jwt.GenerateToken(user.Email, user.Role)
	if err != nil {
		fmt.Println("[Login] token generation failed:", err)
		return "", err
	}
	fmt.Println("[Login] token generated:", token)
	return token, nil
}

func (u *AuthUsecase) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	if user, found := u.cache.Get(email); found {
		fmt.Println("[GetProfile] cache HIT for:", email)
		return user, nil
	}

	fmt.Println("[GetProfile] cache MISS for:", email)
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("[GetProfile] error:", err)
		return nil, err
	}

	u.cache.Set(email, user)
	fmt.Println("[GetProfile] user cached:", email)

	return user, nil
}

func (u *AuthUsecase) UpdateProfile(ctx context.Context, email string, updated *domain.User) (*domain.User, error) {
	user, err := u.repo.UpdateUser(ctx, email, updated)
	if err != nil {
		fmt.Println("[UpdateProfile] error:", err)
		return nil, err
	}

	u.cache.Invalidate(email)
	fmt.Println("[UpdateProfile] cache invalidated for:", email)

	updateEvent := map[string]interface{}{
		"email":     user.Email,
		"role":      user.Role,
		"full_name": user.FullName,
		"house":     user.House,
		"street":    user.Street,
		"apartment": user.Apartment,
		"action":    "profile_updated",
	}

	if eventData, err := json.Marshal(updateEvent); err == nil {
		if err := u.publisher.Publish(ctx, "user.profile_updated", eventData); err != nil {
			fmt.Println("[UpdateProfile] failed to publish user.profile_updated:", err)
		} else {
			fmt.Println("[UpdateProfile] published update event:", string(eventData))
		}
	}

	return user, nil
}
