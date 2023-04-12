package main

import (
	"fmt"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/schedule"
)

func main() {
	config.LoadEnv()
	db.Connect()
	schedule.InitCron()

	app := rest.InitApp()
	app.Listen(fmt.Sprintf(":%s", config.REST_PORT))
}
