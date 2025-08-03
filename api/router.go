package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"os"
	"time"
)

func SetupRouter(dispatcher *jobs.Dispatcher) *gin.Engine {
	router := gin.Default()

	var frontend_url string
	if frontend_url = os.Getenv("FRONTEND_URL"); frontend_url == "" {
		frontend_url = "http://localhost:5173"
	}
	fmt.Printf("Frontend URL: %s", frontend_url)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontend_url},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/upload", handleUpload(dispatcher))               // upload a file
	router.GET("/users/:id/files", handleGetUserFiles(dispatcher)) // get files by user id
	router.GET("/files/:id", handleGetFile(dispatcher))            // get file by file id
	router.DELETE("/files/:id", handleDeleteFile(dispatcher))      // delete file by file id
	// TODO: auth endpoint
	return router
}
