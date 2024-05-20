package model

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	User UserResponse
	jwt.RegisteredClaims
}
