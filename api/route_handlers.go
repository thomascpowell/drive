package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/utils"
	"net/http"
	"strings"
)

func handleAuth(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds models.Credentials
		if err := ctx.ShouldBindJSON(&creds); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.AuthenticateUser,
			Payload: &creds,
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		defer close(job.Done)
		token := <-job.Done
		if token.Err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": token.Err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token.Value})
	}
}

func handleRegister(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds models.Credentials
		if err := ctx.ShouldBindJSON(&creds); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		if strings.TrimSpace(creds.Username) == "" || strings.TrimSpace(creds.Password) == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		hashedPassword, err := utils.HashPassword(creds.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user := models.User{
			Username: creds.Username,
			Password: hashedPassword,
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.RegisterUser,
			Payload: &user,
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		defer close(job.Done)
		result := <-job.Done
		if result.Err != nil { // check for username collision?
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "user registered"})
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
