package middleware

import (
	"api-gateway-service/config"
	"auth-service/pkg/jwt"
	"context"
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
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		email, role, err := jwt.ParseToken(token, cfg.JWTSecret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		ctx = context.WithValue(ctx, UserRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
