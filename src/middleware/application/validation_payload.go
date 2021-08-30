package application

import jwt "github.com/dgrijalva/jwt-go"

type AuthClaims struct {
	ID         string `json:"id"`
	RoleID     int64  `json:"role_id"`
	IsVerified int64  `json:"is_verified"`
	jwt.StandardClaims
}

type PRMNClaims struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	FacebookID string `json:"facebook_id"`
	GoogleID   string `json:"google_id"`
	AppleID    string `json:"apple_id"`
	Status     string `json:"status"`
}
