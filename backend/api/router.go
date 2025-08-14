package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thomascpowell/drive/jobs"
	"os"
	"time"
)

func SetupRouter(dispatcher *jobs.Dispatcher) *gin.Engine {
	router := gin.Default()

	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	router.POST("/upload", handleUpload(dispatcher))          // upload a file
	router.GET("/files", handleGetUserFiles(dispatcher))      // get files by user id
	router.GET("/files/:id", handleGetFile(dispatcher))       // get file by file id
	router.DELETE("/files/:id", handleDeleteFile(dispatcher)) // delete file by file id
	router.POST("/auth", handleAuth(dispatcher))              // authenticate by credentials
	router.POST("/register", handleRegister(dispatcher))      // add a new user
	router.GET("/health", func(c *gin.Context) {              // check connection
		c.JSON(200, gin.H{"status": "ok"})
	})
	return router
}
