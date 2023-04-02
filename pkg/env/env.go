package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv(){
	// open log file
	openEnvFile()
}

func openEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetKey(key string) string {
	return os.Getenv(key)
}

func SetKey(key string, value string) {
	println("setting env: " + key + " to " + value)
	os.Setenv(key, value)
}
