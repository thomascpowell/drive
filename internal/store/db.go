package store

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/thomascpowell/drive/internal/models" 
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

func (s *Store) CreateUser(u *models.User) error {
	return s.DB.Create(u).Error
}

func (s *Store) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) CreateFile(f *models.File) error {
	return s.DB.Create(f).Error
}

func (s *Store) GetFileByID(id uint) (*models.File, error) {
	var file models.File
	err := s.DB.First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (s *Store) GetFilesByUserID(userID uint) ([]models.File, error) {
	var files []models.File
	err := s.DB.Where("uploaded_by = ?", userID).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}




