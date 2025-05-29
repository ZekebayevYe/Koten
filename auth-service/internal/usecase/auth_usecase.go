package usecase

import (
	"auth-service/config"
	"auth-service/internal/domain"
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
}

func NewAuthUsecase(repo domain.UserRepository, publisher domain.EventPublisher, cfg *config.Config) *AuthUsecase {
	return &AuthUsecase{repo: repo, publisher: publisher, config: cfg}
}

func (u *AuthUsecase) Register(ctx context.Context, user *domain.User) (string, error) {
	// 🔥 ЛОГИРОВАНИЕ: Проверяем входящие данные
	fmt.Printf("[Register] INPUT USER: email=%s, full_name=%s, house=%s, street=%s, apartment=%s\n",
		user.Email, user.FullName, user.House, user.Street, user.Apartment)

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

	// 🔥 ДОБАВЛЕНО: Устанавливаем время создания
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// 🔥 ЛОГИРОВАНИЕ: Проверяем данные перед сохранением
	fmt.Printf("[Register] BEFORE SAVE: email=%s, full_name=%s, house=%s, street=%s, apartment=%s, role=%s\n",
		user.Email, user.FullName, user.House, user.Street, user.Apartment, user.Role)

	if err := u.repo.CreateUser(ctx, user); err != nil {
		fmt.Println("[Register] create user error:", err)
		return "", err
	}

	// 🔥 ЛОГИРОВАНИЕ: Проверяем данные после сохранения
	fmt.Printf("[Register] AFTER SAVE: email=%s, full_name=%s, house=%s, street=%s, apartment=%s, role=%s\n",
		user.Email, user.FullName, user.House, user.Street, user.Apartment, user.Role)

	// 🔥 ИСПРАВЛЕНИЕ: Публикуем полные данные пользователя
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
		fmt.Println("[Register] failed to marshal event:", err)
		// Не возвращаем ошибку, так как пользователь уже создан
	} else {
		fmt.Println("[Register] EVENT TO PUBLISH:", string(eventData))
		if err := u.publisher.Publish(ctx, "user.registered", eventData); err != nil {
			fmt.Println("[Register] failed to publish user.registered:", err)
		} else {
			fmt.Println("[Register] ✅ EVENT PUBLISHED SUCCESSFULLY")
		}
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

	// 🔥 ДОПОЛНЕНИЕ: Можно также публиковать событие логина
	loginEvent := map[string]interface{}{
		"email":     user.Email,
		"role":      user.Role,
		"action":    "login",
		"timestamp": fmt.Sprintf("%d", ctx.Value("timestamp")),
	}

	if eventData, err := json.Marshal(loginEvent); err == nil {
		if err := u.publisher.Publish(ctx, "user.login", eventData); err != nil {
			fmt.Println("[Login] failed to publish user.login:", err)
		} else {
			fmt.Println("[Login] published login event:", string(eventData))
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

	// 🔥 ДОПОЛНЕНИЕ: Публикуем событие обновления профиля
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
