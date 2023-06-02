package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	loadDotEnv()

	// rest api
	REST_PORT = os.Getenv("REST_PORT")

	// grpc server
	GRPC_PORT = os.Getenv("GRPC_PORT")

	// database
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PASS = os.Getenv("DB_PASS")
	DB_USER = os.Getenv("DB_USER")

	// image
	IMAGE_HOST_ENDPOINT = os.Getenv("IMAGE_HOST_ENDPOINT")
	IMAGE_EXTENSIONS = strings.Split(os.Getenv("IMAGE_EXTENSIONS"), ",")

	// github api
	GH_TOKEN = os.Getenv("GH_TOKEN")

	// aws
	AWS_ACCESS = os.Getenv("AWS_ACCESS")
	AWS_SECRET = os.Getenv("AWS_SECRET")
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}
}
