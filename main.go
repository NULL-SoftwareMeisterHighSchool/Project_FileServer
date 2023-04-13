package main

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	grpc_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server"
	rest_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest"
)

func main() {
	config.LoadEnv()
	db.Connect()
	go grpc_server.Listen()
	rest_server.Listen()
}
