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

func JWTMiddleware(jwtSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		email, role, err := jwtutil.ParseToken(tokenStr, jwtSecret)
		if err != nil {
			fmt.Println("[Auth Service JWT] Token parsing error:", err)
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyEmail, email)
		ctx = context.WithValue(ctx, ContextKeyRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func EmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value(ContextKeyEmail).(string)
	if !ok || email == "" {
		return "", errors.New("email not found in context")
	}
	return email, nil
}

func RoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value(ContextKeyRole).(string)
	if !ok || role == "" {
		return "", errors.New("role not found in context")
	}
	return role, nil
}
