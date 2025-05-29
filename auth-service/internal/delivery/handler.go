package delivery

import (
	"context"
	"errors"
	"fmt"

	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/internal/usecase"
	pb "auth-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	Usecase *usecase.AuthUsecase
	Cfg     *config.Config
}

// 🔥 ДОБАВЛЕНО: функции для извлечения данных из контекста
func getEmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value("email").(string)
	if !ok || email == "" {
		return "", errors.New("email not found in context")
	}
	return email, nil
}

func getRoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value("role").(string)
	if !ok || role == "" {
		return "", errors.New("role not found in context")
	}
	return role, nil
}

func (h *Handler) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	fmt.Println("[RegisterUser] request:", req)
	user := &domain.User{
		Email:     req.Email,
		Password:  req.Password,
		FullName:  req.FullName,
		House:     req.House,
		Street:    req.Street,
		Apartment: req.Apartment,
	}

	token, err := h.Usecase.Register(ctx, user)
	if err != nil {
		fmt.Println("[RegisterUser] error:", err)
		return nil, err
	}
	fmt.Println("[RegisterUser] token:", token)

	return &pb.AuthResponse{Token: token}, nil
}

func (h *Handler) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	fmt.Println("[LoginUser] request: email =", req.Email, "password =", req.Password)

	token, err := h.Usecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		fmt.Println("[LoginUser] error:", err)
		return nil, err
	}
	fmt.Println("[LoginUser] token:", token)

	return &pb.AuthResponse{Token: token}, nil
}

// 🔥 ИСПРАВЛЕНО: GetMyProfile теперь использует email из контекста
func (h *Handler) GetMyProfile(ctx context.Context, req *pb.GetMyProfileRequest) (*pb.UserProfile, error) {
	// Получаем email из контекста (токен уже проверен в interceptor)
	email, err := getEmailFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %v", err)
	}

	role, err := getRoleFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %v", err)
	}

	fmt.Println("[GetMyProfile] email from context:", email, "role:", role)

	user, err := h.Usecase.GetProfile(ctx, email)
	if err != nil {
		fmt.Println("[GetMyProfile] error:", err)
		return nil, err
	}

	return &pb.UserProfile{
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      role,
		House:     user.House,
		Street:    user.Street,
		Apartment: user.Apartment,
	}, nil
}

// 🔥 ИСПРАВЛЕНО: UpdateMyProfile теперь использует email из контекста
func (h *Handler) UpdateMyProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserProfile, error) {
	fmt.Println("[UpdateMyProfile] request:", req)

	// Получаем email из контекста (токен уже проверен в interceptor)
	email, err := getEmailFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %v", err)
	}

	role, err := getRoleFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %v", err)
	}

	fmt.Println("[UpdateMyProfile] email from context:", email)

	updated := &domain.User{
		FullName:  req.FullName,
		House:     req.House,
		Street:    req.Street,
		Apartment: req.Apartment,
	}

	user, err := h.Usecase.UpdateProfile(ctx, email, updated)
	if err != nil {
		fmt.Println("[UpdateMyProfile] error:", err)
		return nil, err
	}

	return &pb.UserProfile{
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      role,
		House:     user.House,
		Street:    user.Street,
		Apartment: user.Apartment,
	}, nil
}