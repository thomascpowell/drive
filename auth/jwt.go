package auth

import (
	"fmt"
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

func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})
}
