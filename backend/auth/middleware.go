package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing auth header"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}
		tokenStr := parts[1]
		token, err := ParseJWT(tokenStr)
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, exists := claims["sub"].(string); exists {
				ctx.Set("sub", sub)
			}
		}
		ctx.Next()
	}
}


func LoadTokenOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return
		}
		tokenStr := parts[1]
		token, err := ParseJWT(tokenStr)
		if err != nil || !token.Valid {
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, exists := claims["sub"].(string); exists {
				ctx.Set("sub", sub)
			}
		}
		ctx.Next()
	}
}
