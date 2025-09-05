package utils

import (
	"errors"
	"os"
)

func getEnvVariable(key string) (string, error) {
	env := os.Getenv(key)
	if env != "" {
		return env, nil
	}
	return "", errors.New("env variable not set")
}
