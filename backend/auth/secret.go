package auth

import (
	"os"
)

func GetJWTSecret() []byte {
	if os.Getenv("ENVIRONMENT") != "dev" {
		return []byte(os.Getenv("JWT_SECRET"))
	}
	return []byte("test")
}
