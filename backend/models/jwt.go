package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
