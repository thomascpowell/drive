package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/utils"
	"net/http"
	"fmt"
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
		upload, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
			return
		}
		rawID, exists := ctx.Get("sub") // TODO: auth middleware, this should be uint
		userID, ok := rawID.(uint)
		if !exists || !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
			return
		}
		file := &models.File{
			Filename:   upload.Filename,
			Size:       upload.Size,
			UploadedBy: userID,
			Path:       fmt.Sprintf("%d/%s", userID, upload.Filename),
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.Upload,
			Payload: file,
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		result := <-job.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}
		if err := ctx.SaveUploadedFile(upload, file.Path); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "file saved"})
	}
}

func handleGetUserFiles(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get files by user job
		// return []File as json
	}
}

func handleGetFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get file by :id job
		// return file as json
	}
}

func handleDeleteFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// delete file by :id job
		// return status and message
	}
}
