package middleware

import (
	"auth-service/pkg/jwtutil"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type contextKey string

const (
	ContextKeyEmail contextKey = "email"
	ContextKeyRole  contextKey = "role"
)

// JWTMiddleware проверяет JWT токен и сохраняет email и роль в контекст
func JWTMiddleware(jwtSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Auth Service JWT] Processing request to:", r.URL.Path)

		authHeader := r.Header.Get("Authorization")
		fmt.Println("[Auth Service JWT] Authorization header:", authHeader)

		if authHeader == "" {
			fmt.Println("[Auth Service JWT] Missing authorization header")
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("[Auth Service JWT] Invalid header format, parts:", parts)
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		fmt.Println("[Auth Service JWT] Extracted token:", tokenStr)

		// ✅ Исправлено: используем tokenStr и jwtSecret
		email, role, err := jwtutil.ParseToken(tokenStr, jwtSecret)
		if err != nil {
			fmt.Println("[Auth Service JWT] Token parsing error:", err)
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println("[Auth Service JWT] Token validated successfully, email:", email, "role:", role)

		ctx := context.WithValue(r.Context(), ContextKeyEmail, email)
		ctx = context.WithValue(ctx, ContextKeyRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// EmailFromContext извлекает email из контекста
func EmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value(ContextKeyEmail).(string)
	if !ok || email == "" {
		return "", errors.New("email not found in context")
	}
	return email, nil
}

// RoleFromContext извлекает роль из контекста
func RoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value(ContextKeyRole).(string)
	if !ok || role == "" {
		return "", errors.New("role not found in context")
	}
	return role, nil
}
