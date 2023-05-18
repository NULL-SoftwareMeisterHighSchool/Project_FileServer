package grpc_test

import (
	"context"
	"testing"

	user_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	articles_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/articles"
	articles_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func articleBeforeEach(t *testing.T) articles_pb.ArticleServiceClient {
	conn, close := e2e_test.SetupgRPC(t, func(baseServer *grpc.Server) {
		articles_pb.RegisterArticleServiceServer(baseServer, articles_server.ArticleServiceServer{})
	})

	client := articles_pb.NewArticleServiceClient(conn)
	t.Cleanup(close)
	return client
}

func TestCreateArticle(t *testing.T) {
	// given
	e2e_test.Setup(t)
	client := articleBeforeEach(t)

	const userID uint = 1
	user_repo.CreateUserByID(userID)

	type expectation struct {
		input     *articles_pb.CreateArticleRequest
		satusCode codes.Code
	}

	testcases := map[string]expectation{

		"Create article": {
			input: &articles_pb.CreateArticleRequest{
				Type:     articles_pb.ArticleType_TECH,
				AuthorID: uint32(userID),
				Title:    "hi",
			},
			satusCode: codes.OK,
		},
		"Create article with invalid author": {
			input: &articles_pb.CreateArticleRequest{
				Type:     articles_pb.ArticleType_GENERAL,
				AuthorID: 2,
				Title:    "should fail",
			},
			satusCode: codes.Unknown,
		},
	}

	for title, test := range testcases {
		_, err := client.CreateArticle(context.Background(), test.input)
		if err != nil && !errors.ErrStatusMatches(test.satusCode, err) {
			t.Errorf("%s: %v", title, err)
		}
	}

}
