package grpc

import (
	"context"

	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/internal/usecase"
	"auth-service/pkg/jwt"
	pb "auth-service/proto"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	Usecase *usecase.AuthUsecase
	Cfg     *config.Config
}

func (h *Handler) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user := &domain.User{
		Email:     req.Email,
		Password:  req.Password,
		FullName:  req.FullName,
		House:     req.House,
		Street:    req.Street,
		Apartment: req.Apartment,
	}

	token, err := h.Usecase.Register(ctx, user)
	return &pb.AuthResponse{Token: token}, err
}

func (h *Handler) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, err := h.Usecase.Login(ctx, req.Email, req.Password)
	return &pb.AuthResponse{Token: token}, err
}

func (h *Handler) GetMyProfile(ctx context.Context, req *pb.GetMyProfileRequest) (*pb.UserProfile, error) {
	email, role, err := jwt.ParseToken(req.Token, h.Cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	user, err := h.Usecase.GetProfile(ctx, email)
	if err != nil {
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

func (h *Handler) UpdateMyProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserProfile, error) {
	email, _, err := jwt.ParseToken(req.Token, h.Cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	updated := &domain.User{
		FullName:  req.FullName,
		House:     req.House,
		Street:    req.Street,
		Apartment: req.Apartment,
	}

	user, err := h.Usecase.UpdateProfile(ctx, email, updated)
	if err != nil {
		return nil, err
	}

	return &pb.UserProfile{
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      user.Role,
		House:     user.House,
		Street:    user.Street,
		Apartment: user.Apartment,
	}, nil
}
