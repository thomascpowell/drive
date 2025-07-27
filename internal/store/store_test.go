package store

import(
	"testing"
	"github.com/thomascpowell/drive/internal/models"
)


func TestCreateAndGetUser(t *testing.T) {
	s := NewStore(":memory:")

	user := &models.User{Username: "testuser", Password: "hashedpw"}
	if err := s.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	got, err := s.GetUserByUsername("testuser")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if got.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %s", got.Username)
	}
}
