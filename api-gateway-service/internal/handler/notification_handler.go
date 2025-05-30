package handler

import (
	"api-gateway-service/config"
	authpb "api-gateway-service/proto"
	notificationpb "api-gateway-service/proto"
	"context"
	"encoding/json"
	"net/http"
)

type NotificationHandler struct {
	Client notificationpb.NotificationServiceClient
	Auth   authpb.AuthServiceClient
	Cfg    *config.Config
}

func (h *NotificationHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	profile, err := h.Auth.GetMyProfile(context.Background(), &authpb.GetMyProfileRequest{Token: token})
	if err != nil {
		http.Error(w, "auth failed", http.StatusUnauthorized)
		return
	}

	_, err = h.Client.Subscribe(context.Background(), &notificationpb.EmailRequest{
		Email:  profile.Email,
		Street: profile.Street,
		House:  profile.House,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "subscribed"})
}

func (h *NotificationHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	profile, err := h.Auth.GetMyProfile(context.Background(), &authpb.GetMyProfileRequest{Token: token})
	if err != nil {
		http.Error(w, "auth failed", http.StatusUnauthorized)
		return
	}

	_, err = h.Client.Unsubscribe(context.Background(), &notificationpb.EmailRequest{
		Email: profile.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "unsubscribed"})
}
func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	profile, err := h.Auth.GetMyProfile(context.Background(), &authpb.GetMyProfileRequest{Token: token})
	if err != nil {
		http.Error(w, "auth failed", http.StatusUnauthorized)
		return
	}

	if profile.Role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var req notificationpb.Notification
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	_, err = h.Client.CreateNotification(context.Background(), &req)
	if err != nil {
		http.Error(w, "failed to create notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

func (h *NotificationHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	_, err := h.Auth.GetMyProfile(context.Background(), &authpb.GetMyProfileRequest{
		Token: r.Header.Get("Authorization"),
	})
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := h.Client.GetHistory(context.Background(), &notificationpb.Empty{})
	if err != nil {
		http.Error(w, "failed to get history", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Items)
}
