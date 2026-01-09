package utils

import (
	"fmt"
	"os"
)

func GetFilePath() (string, error) {
	path := "./user_files"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetFrontendURL() string {
	if env := os.Getenv("BASE_URL"); env != "" {
		return env
	}
	return "localhost"
}
func GetFrontendURLWithPort() string {
	if env := os.Getenv("BASE_URL"); env != "" {
		return env
	}
	return "localhost:5173"
}

func GetRedisURL() string {
	if env := os.Getenv("REDIS_URL"); env != "" {
		return env
	}
	return "localhost:6379"
}

func GetPostgresDSN() string {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "dev"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "dev"
	}
	name := os.Getenv("POSTGRES_DB")
	if name == "" {
		name = "dev"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
}
