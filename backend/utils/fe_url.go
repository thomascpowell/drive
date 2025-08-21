package utils

import (
	"os"
)

func GetFrontendURL() string {
	env := os.Getenv("FE_URL")
	print(env)
	if env == "" {
		// env is unset -> dev
		return "localhost"
	}
	return env
}
