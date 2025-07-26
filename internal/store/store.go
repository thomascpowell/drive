package store

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	models "github.com/thomascpowell/drive/internal/models" 
)

type Store struct {
	DB *gorm.DB
}

func NewStore(dsn string) *Store {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.File{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return &Store{DB: db}
}
