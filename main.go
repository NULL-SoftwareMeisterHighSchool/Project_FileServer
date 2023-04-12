package main

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/rank"
)

func main() {
	config.LoadEnv()
	db.Connect()
	rank.InitCron()

	app := rest.InitApp()
	app.Listen(":8080")
}
