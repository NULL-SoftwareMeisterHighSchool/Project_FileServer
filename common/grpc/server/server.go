package grpc_server

import (
	"fmt"
	"log"
	"net"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"google.golang.org/grpc"
)

func Listen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %s", config.GRPC_PORT, err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
