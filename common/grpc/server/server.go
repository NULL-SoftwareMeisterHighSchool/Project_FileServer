package grpc_server

import (
	"fmt"
	"log"
	"net"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	articles_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/articles"
	articles_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	comments_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
	rank_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/rank"
	users_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
	"google.golang.org/grpc"
)

func Listen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %s", config.GRPC_PORT, err)
	}

	grpcServer := grpc.NewServer()

	articles_pb.RegisterArticleServiceServer(grpcServer, articles_server.ArticleServiceServer{})
	comments_pb.RegisterCommentServiceServer(grpcServer, comments_pb.UnimplementedCommentServiceServer{})
	users_pb.RegisterUserEventServiceServer(grpcServer, users_pb.UnimplementedUserEventServiceServer{})
	rank_pb.RegisterRankServiceServer(grpcServer, rank_pb.UnimplementedRankServiceServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
