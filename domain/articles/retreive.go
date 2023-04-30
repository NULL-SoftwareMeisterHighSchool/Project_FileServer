package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetArticle(articleID, userID uint) (*pb.GetArticleResponse, error) {
	article, err := article_repo.GetArticleWithBody(articleID, userID)

	if err != nil {
		return nil, status.Error(codes.NotFound, "matching article not found")
	}

	return &pb.GetArticleResponse{
		ArticleID: uint32(article.ID),
		AuthorID:  uint32(article.AuthorID),
		Type:      pb.ArticleType(article.Type),
		Title:     article.Title,
		CreatedAt: timestamppb.New(article.CreatedAt),
		UpdatedAt: timestamppb.New(article.UpdatedAt),
		Views:     article.Views,
		Body:      string(article.Body.Text),
		Likes:     uint32(article.Likes),
		IsPrivate: article.IsPrivate,
		IsLiked:   article.IsLiked,
	}, nil
}
