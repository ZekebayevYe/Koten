package handler

import (
	"api-gateway-service/proto"
	"context"
	"encoding/json"
	"net/http"
)

type NewsHandler struct {
	Client proto.NewsServiceClient
}

func (h *NewsHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Client.GetNews(context.Background(), &proto.Empty{})
	if err != nil {
		http.Error(w, "Failed to fetch news", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
