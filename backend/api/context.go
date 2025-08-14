package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserID(ctx *gin.Context) (uint, bool) {
	rawID, exists := ctx.Get("sub")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return 0, false
	}
	userID, ok := rawID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return 0, false
	}
	return userID, true
}

func GetSlug(ctx *gin.Context, param string) (uint, bool) {
	idStr := ctx.Param(param)
	idUint64, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return 0, false
	}
	return uint(idUint64), true
}
