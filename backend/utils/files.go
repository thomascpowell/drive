package utils

import (
	"os"
	"path/filepath"
)

func GetFilePath() (string, error) {
	path := filepath.Join(".", "data", "user_files") // TODO: env this?
	println(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

