package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email, role, secret, expires string) (string, error) {
	duration, err := time.ParseDuration(expires)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr, secret string) (email string, role string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}

	email, _ = claims["email"].(string)
	role, _ = claims["role"].(string)

	return email, role, nil
}
