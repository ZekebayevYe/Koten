package jwtutil

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// ParseToken парсит JWT токен и возвращает email и роль
func ParseToken(tokenStr, secret string) (string, string, error) {
	fmt.Println("[JWTUtil] Received token:", tokenStr)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("[JWTUtil] Parse error:", err)
		return "", "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println("[JWTUtil] Parsed email:", claims.Email, "role:", claims.Role)
		return claims.Email, claims.Role, nil
	}

	return "", "", fmt.Errorf("invalid token")
}
