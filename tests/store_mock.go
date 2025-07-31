package tests

import (
	"github.com/thomascpowell/drive/models"
)

type MockStore struct {
	CreateFileFunc       func(*models.File) error
	GetFilesByUserIDFunc func(uint) ([]models.File, error)
	GetFileByIDFunc      func(uint) (*models.File, error)
	DeleteFileByIDFunc   func(uint) error
}

func (m *MockStore) CreateUser(u *models.User) error {
	return nil
}
func (m *MockStore) GetUserByUsername(username string) (*models.User, error) {
	return nil, nil
}
func (m *MockStore) CreateFile(f *models.File) error {
	return m.CreateFileFunc(f)
}
func (m *MockStore) GetFileByID(id uint) (*models.File, error) {
	return m.GetFileByIDFunc(id)
}
func (m *MockStore) GetFilesByUserID(userID uint) ([]models.File, error) {
	return m.GetFilesByUserIDFunc(userID)
}
func (m *MockStore) DeleteFileByID(id uint) error {
	return m.DeleteFileByIDFunc(id)
}



