package model

import (
    "github.com/dgrijalva/jwt-go"
)

// ClaimsはJWTトークンのクレーム
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}