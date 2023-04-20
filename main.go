package main

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	grpc_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server"
	rest_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest/server"
)

func main() {
	config.LoadEnv()
	database.Connect()
	go grpc_server.Listen()
	rest_server.Listen()
}
