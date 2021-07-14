package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load environments variables.")
	}

	return os.Getenv(key)
}
