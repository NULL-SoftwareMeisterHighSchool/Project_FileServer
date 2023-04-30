package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetArticle(articleID, userID uint) (*pb.GetArticleResponse, error) {

	pInfo, err := article_repo.GetArticlePermissionInfoByID(articleID)

	if err != nil {
		return nil, article_errors.ErrNotFound
	}
	if pInfo.IsPrivate && pInfo.AuthorID != userID {
		return nil, article_errors.ErrPermissionDenied
	}

	article, _ := article_repo.GetArticleWithBody(articleID, userID)

	go article_repo.IncreaseViewCount(articleID)

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
