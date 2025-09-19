package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/utils"
)

func handleGetSharedFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key, ok := GetSlug(ctx, "key")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid key"})
			return
		}

		// holy type conversions
		str_id, err := dispatcher.Redis.Get(fmt.Sprintf("%d", key))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid key"})
		}
		int_id, err := strconv.Atoi(str_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		}
		uint_id := uint(int_id)
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetFile,
			Payload: models.NewGetFilePayload(uint_id),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)

		result := <-job.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
		}
		fileptr, ok := result.Value.(*models.File)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid job result type"})
			return
		}
		file := *fileptr
		basePath, err := utils.GetFilePath()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filePath := filepath.Join(basePath, file.Path)
		ctx.Header("Content-Disposition", `attachment; filename="`+file.Filename+`"`)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(filePath)
	}
}

func handleGetShareLink(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ShareRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		raw, err := strconv.Atoi(req.FileID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		}
		fileID := uint(raw)
		raw, err = strconv.Atoi(req.TTL)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		}
		ttl := uint(raw)
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetShareLink,
			Payload: models.NewGetShareLinkPayload(fileID, ttl),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		link := <-job.Done
		if link.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": link.Err.Error()})
			return
		}
		res := utils.GetFrontendURL() + "/share/" + link.Value.(string)
		ctx.JSON(http.StatusOK, gin.H{"message": res})
	}
}

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
		userID, ok := ctx.Get("sub")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		file := models.File{
			Filename:   upload.Filename,
			Size:       upload.Size,
			UploadedBy: userID.(uint),
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

		basePath, err := utils.GetFilePath()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}
		fullPath := filepath.Join(basePath, file.Path)

		if err := ctx.SaveUploadedFile(upload, fullPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "file saved"})
	}
}

func handleGetUserFiles(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, ok := ctx.Get("sub")
		if !ok {
			return
		}
		job := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetUserFiles,
			Payload: models.NewGetUserFilesPayload(userID.(uint)),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(job)
		result := <-job.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": result.Value})

	}
}

func handleGetFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, ok := ctx.Get("sub")
		if !ok {
			return
		}
		fileID, ok := GetSlug(ctx, "id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
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
		fileptr, ok := result.Value.(*models.File)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid job result type"})
			return
		}
		file := *fileptr
		if file.UploadedBy != userID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user does not have access to this file"})
		}
		basePath, err := utils.GetFilePath()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filePath := filepath.Join(basePath, file.Path)
		ctx.Header("Content-Disposition", `attachment; filename="`+file.Filename+`"`)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(filePath)
	}
}

func handleHealth(ctx *gin.Context) {
	id, exists := ctx.Get("sub")
	message := "ok"
	if exists {
		message += ", has token for id: " + fmt.Sprint(id)
	}
	ctx.JSON(200, gin.H{"message": message})
}

func handleLogout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "", true, true)
	ctx.JSON(200, gin.H{"message": "log out successful"})
}

func handleDeleteFile(dispatcher *jobs.Dispatcher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, ok := ctx.Get("sub")
		if !ok {
			return
		}
		fileID, ok := GetSlug(ctx, "id")
		if !ok {
			return
		}
		getFileJob := &models.Job{
			ID:      utils.UUID(),
			Type:    models.GetFile,
			Payload: models.NewGetFilePayload(fileID),
			Done:    make(chan models.Result, 1),
		}
		dispatcher.Dispatch(getFileJob)
		result := <-getFileJob.Done
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}
		fileptr, ok := result.Value.(*models.File)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid job result type"})
			return
		}
		file := *fileptr
		if file.UploadedBy != userID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user may not delete this file"})
		}
		basePath, err := utils.GetFilePath()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filePath := filepath.Join(basePath, file.Path)
		err = os.Remove(filePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		deleteFileJob := &models.Job{
			ID:      utils.UUID(),
			Type:    models.DeleteFile,
			Payload: models.NewDeleteFilePayload(userID.(uint), fileID),
			// NOTE: userID is currently not checked in the job
			// It is checked in this handler for now
			Done: make(chan models.Result, 1),
		}
		dispatcher.Dispatch(deleteFileJob)
		result2 := <-deleteFileJob.Done
		if result2.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
	}
}
