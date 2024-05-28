package web

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	User UserResponse `json:"user,omitempty"`
	jwt.RegisteredClaims
}
