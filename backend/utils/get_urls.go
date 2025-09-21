package utils

import (
	"os"
)

func GetFrontendURL() string {
	env := os.Getenv("FE_URL")
		// env is unset -> no docker -> dev
		// env is set -> docker -> use the set url
	if env == "" {
		return "localhost"
	}
	return env
}

func GetFrontendURLWithPort() string {
	env := os.Getenv("FE_URL")
	if env == "" {
		return "localhost:5173"
	}
	// basically like the above
	// no need to add port for prod obv
	return env
}

// used for redis (obviously)
func GetRedisURL() string {
	env := os.Getenv("ENVIRONMENT")
		// env is unset -> no docker -> dev
		// env is set -> docker -> use docker relative
	if env == "" {
		return "localhost"
	}
	return "redis"
}

