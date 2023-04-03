package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	loadDotEnv()

	// database
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PASS = os.Getenv("DB_PASS")
	DB_USER = os.Getenv("DB_USER")

	// image
	IMAGE_HOST_ENDPOINT = os.Getenv("IMAGE_HOST_ENDPOINT")
	IMAGE_EXTENSIONS = strings.Split(os.Getenv("IMAGE_EXTENSIONS"), ",")

	// websocket
	WS_REQUEST_ORIGIN = os.Getenv("WS_REQUEST_ORIGIN")
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}
}
