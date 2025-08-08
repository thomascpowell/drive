package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
)

func handleAuth(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func handleRegister(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func handleUpload(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func handleGetUserFiles(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func handleGetFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func handleDeleteFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
