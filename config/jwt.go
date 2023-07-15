package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("asafeewtery346366346436erygfdhdfh")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
