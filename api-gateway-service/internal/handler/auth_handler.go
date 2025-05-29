package handler

import (
	"api-gateway-service/config"
	pb "api-gateway-service/proto"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	Client pb.AuthServiceClient
	Cfg    *config.Config
}

// 🔥 УЛУЧШЕНО: функция для передачи токена в gRPC metadata
func forwardAuthToken(r *http.Request, ctx context.Context) context.Context {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Логируем для отладки
		fmt.Println("[forwardAuthToken] Authorization header:", authHeader)

		// Передаем токен в gRPC metadata
		md := metadata.Pairs("authorization", authHeader)
		ctx = metadata.NewOutgoingContext(ctx, md)
	} else {
		fmt.Println("[forwardAuthToken] No Authorization header found")
	}
	return ctx
}

// 🔥 НОВАЯ: функция для обработки ошибок
func handleGRPCError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	st, ok := status.FromError(err)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	switch st.Code().String() {
	case "Unauthenticated":
		http.Error(w, st.Message(), http.StatusUnauthorized)
	case "InvalidArgument":
		http.Error(w, st.Message(), http.StatusBadRequest)
	case "NotFound":
		http.Error(w, st.Message(), http.StatusNotFound)
	case "AlreadyExists":
		http.Error(w, st.Message(), http.StatusConflict)
	default:
		http.Error(w, st.Message(), http.StatusInternalServerError)
	}
}

// 🔥 УЛУЧШЕНО: функция для установки JSON заголовков
func setJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Register] Processing registration request")

	// Устанавливаем заголовки
	setJSONHeaders(w)

	var req pb.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("[Register] JSON decode error:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Логируем запрос (без пароля)
	fmt.Printf("[Register] Request: email=%s, full_name=%s, house=%s, street=%s, apartment=%s\n",
		req.Email, req.FullName, req.House, req.Street, req.Apartment)

	resp, err := h.Client.RegisterUser(context.Background(), &req)
	if err != nil {
		fmt.Println("[Register] gRPC error:", err)
		handleGRPCError(w, err)
		return
	}

	fmt.Println("[Register] Success, token generated")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"token":   resp.Token,
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Login] Processing login request")

	setJSONHeaders(w)

	var req pb.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("[Login] JSON decode error:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("[Login] Request: email=%s\n", req.Email)

	res, err := h.Client.LoginUser(context.Background(), &req)
	if err != nil {
		fmt.Println("[Login] gRPC error:", err)
		handleGRPCError(w, err)
		return
	}

	fmt.Println("[Login] Success, token generated")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token":   res.Token,
		"message": "Login successful",
	})
}

// 🔥 ИСПРАВЛЕНО: GetProfile с лучшей обработкой ошибок
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GetProfile] Processing get profile request")

	setJSONHeaders(w)

	ctx := context.Background()
	ctx = forwardAuthToken(r, ctx)

	resp, err := h.Client.GetMyProfile(ctx, &pb.GetMyProfileRequest{})
	if err != nil {
		fmt.Println("[GetProfile] gRPC error:", err)
		handleGRPCError(w, err)
		return
	}

	fmt.Printf("[GetProfile] Success: email=%s, role=%s\n", resp.Email, resp.Role)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// 🔥 ИСПРАВЛЕНО: UpdateProfile с лучшей обработкой
func (h *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[UpdateProfile] Processing update profile request")

	setJSONHeaders(w)

	ctx := context.Background()
	ctx = forwardAuthToken(r, ctx)

	var req pb.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("[UpdateProfile] JSON decode error:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("[UpdateProfile] Request: full_name=%s, house=%s, street=%s, apartment=%s\n",
		req.FullName, req.House, req.Street, req.Apartment)

	res, err := h.Client.UpdateMyProfile(ctx, &req)
	if err != nil {
		fmt.Println("[UpdateProfile] gRPC error:", err)
		handleGRPCError(w, err)
		return
	}

	fmt.Printf("[UpdateProfile] Success: email=%s\n", res.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    res,
		"message": "Profile updated successfully",
	})
}

// 🔥 НОВОЕ: Middleware для логирования запросов
func (h *AuthHandler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", r.Method, r.URL.Path, r.Header.Get("Authorization"))
		next.ServeHTTP(w, r)
	})
}

// 🔥 НОВОЕ: OPTIONS handler для CORS
func (h *AuthHandler) HandleOptions(w http.ResponseWriter, r *http.Request) {
	setJSONHeaders(w)
	w.WriteHeader(http.StatusOK)
}
