package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetArticle(articleID, userID uint) (*pb.GetArticleResponse, error) {

	isAuthor, err := article_utils.CheckPrivateAndExists(userID, articleID)
	if err != nil {
		return nil, err
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
		Body:      string(article.Text),
		Likes:     uint32(article.Likes),
		IsPrivate: article.IsPrivate,
		IsLiked:   article.IsLiked,
		IsAuthor:  isAuthor,
		Comments:  uint32(article.Comments),
	}, nil
}
