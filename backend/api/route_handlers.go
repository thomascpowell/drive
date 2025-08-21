package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/utils"
	"net/http"
	"path/filepath"
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
			Payload: models.NewAuthenticateUserPayload(creds.Username, creds.Password),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		token := <-job.Done
		if token.Err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": token.Err.Error()})
			return
		}
		day := 86400
		secure := false
		frontend := utils.GetFrontendURL()
		httpOnly := true
		ctx.SetCookie("jwt", token.Value.(string), day, "/", frontend, secure, httpOnly)
		ctx.JSON(http.StatusOK, gin.H{"message": "log in successful"})
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
			Payload: models.NewRegisterUserPayload(user),
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
		userID, ok := GetUserID(ctx)
		if !ok {
			return
		}
		file := models.File{
			Filename:   upload.Filename,
			Size:       upload.Size,
			UploadedBy: userID,
			Path:       fmt.Sprintf("%d/%s", userID, upload.Filename),
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.Upload,
			Payload: models.NewUploadPayload(file),
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
		userID, ok := GetUserID(ctx)
		if !ok {
			return
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetUserFiles,
			Payload: models.NewGetUserFilesPayload(userID),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		result := <-job.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": result})

	}
}

func handleGetFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, ok := GetUserID(ctx)
		if !ok {
			return
		}
		fileID, ok := GetSlug(ctx, "id")
		if !ok {
			return
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetFile,
			Payload: models.NewGetFilePayload(fileID),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		result := <-job.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
		}
		file, ok := result.Value.(models.File)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid job result type"})
			return
		}
		if file.UploadedBy != userID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user does not have access to this file"})
		}
		basePath, err := utils.GetFilePath()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filePath := filepath.Join(basePath, file.Path)
		ctx.File(filePath)
	}
}

func handleHealth(ctx *gin.Context) {
	id, exists := ctx.Get("sub")
	message := "ok"
	if exists {
		message += ", has token for id: " + fmt.Sprint(id)
	}
	fmt.Print(id)
	ctx.JSON(200, gin.H{"message": message})
}

func handleDeleteFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO
	}
}
