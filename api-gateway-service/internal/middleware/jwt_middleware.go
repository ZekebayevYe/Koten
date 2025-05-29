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

func AuthMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		email, role, err := jwtutil.ParseToken(tokenStr, cfg.JWTSecret)
		if err != nil {
			fmt.Println("[Gateway Auth] Token parsing error:", err)
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		ctx = context.WithValue(ctx, UserRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
