package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TOKEN_DURATION = 24 * time.Hour

func GenerateJWT(id uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(TOKEN_DURATION).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}
