package handler

import (
	"api-gateway-service/config"
	pb "api-gateway-service/proto"
	"context"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Client pb.AuthServiceClient
	Cfg    *config.Config
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
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.Client.LoginUser(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	resp, err := h.Client.GetMyProfile(context.Background(), &pb.GetMyProfileRequest{Token: token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

