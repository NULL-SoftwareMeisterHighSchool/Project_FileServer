package articles

import (
	"time"

	repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ListArticles(
	offset uint,
	amount uint,
	order repo.ArticleOrder,
	articleType repo.ArticleTypeOption,
	authorID uint,
	userID uint,
	duration *pb.Duration,
	query string,
) (*pb.ListArticleResponse, error) {

	falsyP := false
	start, end := convertDurationOrDefault(duration)

	articles, totalCnt := repo.ListArticles(
		offset, amount, order, articleType,
		authorID, userID, &falsyP, start, end, query,
	)

	return &pb.ListArticleResponse{
		Articles:   convertToResponseElems(articles),
		TotalCount: uint32(totalCnt),
	}, nil
}

func ListArticlesByAuthor(
	offset uint,
	amount uint,
	order repo.ArticleOrder,
	articleType repo.ArticleTypeOption,
	authorID uint,
	userID uint,
	duration *pb.Duration,
	query string,
	isPrivate *bool,
) (*pb.ListArticleResponse, error) {

	start, end := convertDurationOrDefault(duration)

	if userID != authorID {
		if isPrivate == nil {
			*isPrivate = false
		} else if *isPrivate {
			return nil, article_errors.ErrPermissionDenied
		}
	}

	articles, totalCnt := repo.ListArticles(
		offset, amount, order, articleType,
		authorID, userID, isPrivate, start, end, query,
	)

	return &pb.ListArticleResponse{
		Articles:   convertToResponseElems(articles),
		TotalCount: uint32(totalCnt),
	}, nil
}

func convertDurationOrDefault(duration *pb.Duration) (start time.Time, end time.Time) {
	if duration != nil {
		start = time.Unix(duration.GetStart().GetSeconds(), 0)
		end = time.Unix(duration.GetEnd().GetSeconds(), 0)
	} else {
		now := time.Now()
		start = now.AddDate(0, -1, 0) // a month ago
		end = now
	}
	return
}

func convertToResponseElems(articles []*repo.ListArticleElemWithLikes) []*pb.ListArticleElement {
	var elems []*pb.ListArticleElement

	for _, article := range articles {
		elems = append(elems, &pb.ListArticleElement{
			ArticleID: uint32(article.ID),
			AuthorID:  uint32(article.AuthorID),
			Title:     article.Title,
			Summary:   article.Summary,
			Type:      pb.ArticleType(article.Type),
			IsPrivate: article.IsPrivate,
			CreatedAt: timestamppb.New(article.CreatedAt),
			Views:     article.Views,
			Likes:     uint32(article.Likes),
			Thumbnail: article.Thumbnail,
			IsAuthor:  article.IsAuthor,
			Comments:  uint32(article.Comments),
		})
	}

	return elems
}
