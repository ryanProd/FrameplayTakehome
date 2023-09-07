package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Used to access the .env variables
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)
}
