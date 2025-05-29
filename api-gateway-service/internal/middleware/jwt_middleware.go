package middleware

import (
	"api-gateway-service/config"
	"auth-service/pkg/jwtutil"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type key string

const (
	UserEmailKey key = "user_email"
	UserRoleKey  key = "user_role"
)

// AuthMiddleware проверяет JWT токен и сохраняет email и роль в контекст
func AuthMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Gateway Auth] Processing request to:", r.URL.Path)

		authHeader := r.Header.Get("Authorization")
		fmt.Println("[Gateway Auth] Authorization header:", authHeader)

		if authHeader == "" {
			fmt.Println("[Gateway Auth] Missing authorization header")
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("[Gateway Auth] Invalid authorization header format")
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		fmt.Println("[Gateway Auth] Extracted token:", tokenStr)

		// ✅ Исправлено: используем tokenStr и cfg.JWTSecret
		email, role, err := jwtutil.ParseToken(tokenStr, cfg.JWTSecret)
		if err != nil {
			fmt.Println("[Gateway Auth] Token parsing error:", err)
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println("[Gateway Auth] Token validated successfully, email:", email, "role:", role)

		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		ctx = context.WithValue(ctx, UserRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
