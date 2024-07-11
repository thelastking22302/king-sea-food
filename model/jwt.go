package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID string `json:"user_id"`
	Role   string `json:"-"`
	jwt.StandardClaims
}
