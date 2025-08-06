package tests

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/thomascpowell/drive/auth"
	"github.com/thomascpowell/drive/utils"
	"testing"
	"time"
)

func generateJWT(t *testing.T) string {
	userID := uint(0)
	validJWT, err := auth.GenerateJWT(userID)
	if err != nil {
		t.Fatal(err)
	}
	return validJWT
}

func TestValidToken(t *testing.T) {
	validJWT := generateJWT(t)
	parsed, err := auth.ParseJWT(validJWT)
	if err != nil {
		t.Fatalf("error parsing jwt: %v", err)
	}
	claims := parsed.Claims.(jwt.MapClaims)
	utils.Expect(t, int(claims["sub"].(float64)), 0, "subject claim")
}

func TestWrongAlgorithm(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"sub": 0,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	badToken, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("failed to sign with none: %v", err)
	}
	_, err = auth.ParseJWT(badToken)
	if err == nil {
		t.Fatal("expected error for invalid signing method")
	}
}

func TestTamperedToken(t *testing.T) {
	validJWT := generateJWT(t)
	invalid := validJWT + "tampering"
	_, err := auth.ParseJWT(invalid)
	utils.Expect(t, err != nil, true, "error")
}
