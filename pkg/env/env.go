package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	// open log file
	openEnvFile()
}

func openEnvFile() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("Error loading .env file: %s\n", err)
		os.Exit(1)
	}
}

func GetKey(key string) string {
	return os.Getenv(key)
}

func SetKey(key string, value string) {
	println("setting env: " + key + " to " + value)
	err := os.Setenv(key, value)
	if err != nil {
		log.Printf("error setting env key: %v", err)
	}
}
