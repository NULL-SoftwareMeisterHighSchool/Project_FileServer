package articles

import (
	"time"

	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ListArticles(
	offset uint,
	amount uint,
	order repo.ArticleOrder,
	articleType repo.ArticleTypeOption,
	authorID uint,
	isPrivate *bool,
	duration *pb.Duration,
	query string,
) (*pb.ListArticleResponse, error) {

	var start time.Time
	var end time.Time

	if duration != nil {
		start = time.Unix(duration.GetStart().GetSeconds(), 0)
		end = time.Unix(duration.GetEnd().GetSeconds(), 0)
	} else {
		now := time.Now()
		start = now.AddDate(0, 1, 0) // a month ago
		end = now
	}

	articles := article_repo.ListArticles(
		offset, amount, order, articleType,
		authorID, isPrivate, start, end, query,
	)

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
			UpdatedAt: timestamppb.New(article.UpdatedAt),
			Views:     article.Views,
			Likes:     uint32(article.Likes),
		})
	}

	return &pb.ListArticleResponse{
		Articles: elems,
	}, nil
}
