package articles_server

import (
	"context"

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

func (ArticleServiceServer) ListArticle(context.Context, *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {

}

func (ArticleServiceServer) GetArticle(context.Context, *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {

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
