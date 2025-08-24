package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSlug(ctx *gin.Context, param string) (uint, bool) {
	idStr := ctx.Param(param)
	idUint64, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return 0, false
	}
	return uint(idUint64), true
}
