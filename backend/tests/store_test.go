package tests

import (
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/store"
	"testing"
	"fmt"
)

func TestDBOperations(t *testing.T) {
	s := store.NewStore(":memory:")

	testUser := &models.User{
		Username: "testuser",
		Password: "pw",
	}
	if err := s.CreateUser(testUser); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	user, err := s.GetUserByUsername("testuser")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %s", user.Username)
	}
	testUserID := user.ID

	testFile := &models.File{
		Filename:   "testfile.txt",
		Path:       fmt.Sprintf("%d/testfile.txt", user.ID),
		Size:       1,
		UploadedBy: testUserID,
	}
	if err := s.CreateFile(testFile); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	file, err := s.GetFileByID(testFile.ID)
	if err != nil {
		t.Fatalf("failed to get file by id: %v", err)
	}
	if file.UploadedBy != testFile.UploadedBy || file.ID != testFile.ID {
		t.Fatalf("incorrect file returned")
	}

	files, err := s.GetFilesByUserID(testUserID)
	if err != nil {
		t.Fatalf("failed to get files by user id: %v", err)
	}
	if len(files) != 1 {
		t.Errorf("expected 1 file, got %d", len(files))
	}

	if err := s.DeleteFileByID(testFile.ID); err != nil {
		t.Fatalf("failed to delete file: %v", err)
	}
	if _, err := s.GetFileByID(testFile.ID); err == nil {
		t.Errorf("expected error accessing deleted file, got none")
	}
}
