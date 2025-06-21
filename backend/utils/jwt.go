package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sdafbja/telecom-shop/models"
)

// Lấy secret từ env
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT tạo token dựa trên JWTClaims
func GenerateJWT(userID uint, role string) (string, error) {
	claims := models.JWTClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseJWT parse token và trả về JWTClaims
func ParseJWT(tokenString string) (*models.JWTClaims, error) {
	claims := &models.JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
