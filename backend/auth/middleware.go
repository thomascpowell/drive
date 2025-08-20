package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: reimplement
		ctx.Next()
	}
}

func LoadTokenOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenRaw, err := ctx.Cookie("jwt")
		if err != nil {
			ctx.Next()
			return
		}
		token, err := ParseJWT(tokenRaw)
		if err != nil || !token.Valid {
			ctx.Next()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, exists := claims["sub"].(float64); exists {
				ctx.Set("sub", uint(sub))
			}
		}
		ctx.Next()
	}
}
