package articles_server

import (
	"context"

	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles"
	"github.com/golang/protobuf/ptypes/empty"
)

type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
}

func (ArticleServiceServer) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*empty.Empty, error) {
	if err := articles.CreateArticle(uint(request.GetAuthorID()), request.GetTitle(), request.GetType()); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (ArticleServiceServer) ListArticle(ctx context.Context, request *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {
	return articles.ListArticles(
		uint(request.GetOffset()),
		uint(request.GetAmount()),
		article_repo.ArticleOrder(request.GetOrder()),
		article_repo.ArticleTypeOption(request.GetType()),
		uint(request.GetAuthorID()),
		request.IsPrivate,
		request.GetDuration(),
		request.GetQuery(),
	)
}

func (ArticleServiceServer) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	return articles.GetArticle(uint(request.GetArticleID()), uint(request.GetUserID()))
}

func (ArticleServiceServer) UpdateArticleBody(context.Context, *pb.UpdateArticleBodyRequest) (*empty.Empty, error) {

}

func (ArticleServiceServer) UpdateArticleTitle(context.Context, *pb.UpdateArticleTitleRequest) (*empty.Empty, error) {

}

func (ArticleServiceServer) DeleteArticle(context.Context, *pb.DeleteArticleRequest) (*empty.Empty, error) {

}

func (ArticleServiceServer) SetArticleVisibility(context.Context, *pb.SetArticleVisibilityRequest) (*empty.Empty, error) {

}

func (ArticleServiceServer) ToggleArticleLike(context.Context, *pb.ToggleArticleLikeRequest) (*empty.Empty, error) {

}
