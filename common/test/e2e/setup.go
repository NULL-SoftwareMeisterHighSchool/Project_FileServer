package e2e_test

import (
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	"github.com/joho/godotenv"
)

func Setup(t *testing.T) {
	if err := godotenv.Load("../.test.env"); err != nil {
		t.Fatalf("cannot load .test.env: %v", err)
	}
	config.LoadEnv()
	database.Connect()
}
