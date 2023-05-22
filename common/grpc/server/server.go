package grpc_server

import (
	"fmt"
	"log"
	"net"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	articles_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/articles"
	comments_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/comments"
	articles_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	comments_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
	users_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
	users_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/users"
	"google.golang.org/grpc"
)

func Listen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %s", config.GRPC_PORT, err)
	}

	grpcServer := grpc.NewServer()

	articles_pb.RegisterArticleServiceServer(grpcServer, articles_server.ArticleServiceServer{})
	comments_pb.RegisterCommentServiceServer(grpcServer, comments_server.CommentServiceServer{})
	users_pb.RegisterUserEventServiceServer(grpcServer, users_server.UserServiceServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
