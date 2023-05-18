package grpc_article_test

import (
	"testing"

	articles_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/articles"
	articles_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
	"google.golang.org/grpc"
)

func articleBeforeEach(t *testing.T) articles_pb.ArticleServiceClient {
	conn, close := e2e_test.SetupgRPC(t, func(baseServer *grpc.Server) {
		articles_pb.RegisterArticleServiceServer(baseServer, articles_server.ArticleServiceServer{})
	})

	client := articles_pb.NewArticleServiceClient(conn)
	t.Cleanup(close)
	return client
}
