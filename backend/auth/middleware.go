package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getIdFromToken(ctx)
		if id == 0 || err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Set("sub", id)
		ctx.Next()
	}
}

func LoadIdOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getIdFromToken(ctx)
		if id == 0 || err != nil {
			ctx.Next()
			return
		}
		ctx.Set("sub", id)
		ctx.Next()
	}
}

func getIdFromToken(ctx *gin.Context) (uint, error) {
	tokenRaw, err := ctx.Cookie("jwt")
	if err != nil {
		return 0, err
	}
	token, err := ParseJWT(tokenRaw)
	if err != nil || !token.Valid {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, exists := claims["sub"].(float64); exists {
			return uint(sub), nil
		}
	}
	return 0, nil
}
