package handler

import (
	"api-gateway-service/proto"
	"context"
	"encoding/json"
	"net/http"
)

type ReportHandler struct {
	Client proto.ReportServiceClient
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var req proto.CreateReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.Client.CreateReport(context.Background(), &req)
	if err != nil {
		http.Error(w, "Failed to create report", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *ReportHandler) GetReportsByUser(w http.ResponseWriter, r *http.Request) {
	var req proto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.Client.GetReportsByUser(context.Background(), &req)
	if err != nil {
		http.Error(w, "Failed to get reports", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *ReportHandler) EditReport(w http.ResponseWriter, r *http.Request) {
	var req proto.EditReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.Client.EditReport(context.Background(), &req)
	if err != nil {
		http.Error(w, "Failed to edit report", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
