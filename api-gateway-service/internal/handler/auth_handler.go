package handler

import (
	"api-gateway-service/config"
	pb "api-gateway-service/proto"
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/metadata"
)

type AuthHandler struct {
	Client pb.AuthServiceClient
	Cfg    *config.Config
}

// 🔥 ДОБАВЛЕНО: функция для передачи токена в gRPC metadata
func forwardAuthToken(r *http.Request, ctx context.Context) context.Context {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Передаем токен в gRPC metadata
		md := metadata.Pairs("authorization", authHeader)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req pb.RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.Client.RegisterUser(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req pb.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Client.LoginUser(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token": res.Token,
	})
}

// 🔥 ИСПРАВЛЕНО: GetProfile с передачей токена через metadata
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Передаем токен в gRPC metadata
	ctx = forwardAuthToken(r, ctx)

	resp, err := h.Client.GetMyProfile(ctx, &pb.GetMyProfileRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

// 🔥 ИСПРАВЛЕНО: UpdateProfile с передачей токена через metadata
func (h *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Передаем токен в gRPC metadata
	ctx = forwardAuthToken(r, ctx)

	var req pb.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Убираем Token из request, так как он теперь в metadata
	res, err := h.Client.UpdateMyProfile(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
