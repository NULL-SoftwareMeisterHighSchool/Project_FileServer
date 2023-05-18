package grpc_test

import (
	"context"
	"testing"

	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
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

func TestDeleteArticle(t *testing.T) {
	// given
	e2e_test.Setup(t)
	client := articleBeforeEach(t)

	const (
		userID         uint = 1
		weirdUserID    uint = 2
		weirdArticleID uint = 500
	)

	user_repo.CreateUserByID(userID)
	articleID, _ := article_repo.CreateArticle(userID, "", article_entity.TYPE_GENERAL)

	type expectation struct {
		input     *articles_pb.DeleteArticleRequest
		satusCode codes.Code
	}

	testcases := map[string]expectation{

		"Delete without permission": {
			input: &articles_pb.DeleteArticleRequest{
				ArticleID: uint32(articleID),
				UserID:    uint32(weirdUserID),
			},
			satusCode: codes.PermissionDenied,
		},
		"Delete article": {
			input: &articles_pb.DeleteArticleRequest{
				ArticleID: uint32(articleID),
				UserID:    uint32(userID),
			},
			satusCode: codes.OK,
		},
		"Delete inexsisting article": {
			input: &articles_pb.DeleteArticleRequest{
				ArticleID: uint32(weirdArticleID),
				UserID:    uint32(userID),
			},
			satusCode: codes.NotFound,
		},
	}

	for title, test := range testcases {
		_, err := client.DeleteArticle(context.Background(), test.input)
		if err != nil && !errors.ErrStatusMatches(test.satusCode, err) {
			t.Errorf("%s: %v", title, err)
		}
	}

}
