package articles_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
)

type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
}

func (ArticleServiceServer) CreateArticle(context.Context, *pb.CreateArticleRequest) (*pb.StatusResponse, error) {

}

func (ArticleServiceServer) ListArticle(context.Context, *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {

}

func (ArticleServiceServer) GetArticle(context.Context, *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {

}

func (ArticleServiceServer) UpdateArticleBody(context.Context, *pb.UpdateArticleBodyRequest) (*pb.StatusResponse, error) {

}

func (ArticleServiceServer) UpdateArticleTitle(context.Context, *pb.UpdateArticleTitleRequest) (*pb.StatusResponse, error) {

}

func (ArticleServiceServer) DeleteArticle(context.Context, *pb.DeleteArticleRequest) (*pb.StatusResponse, error) {

}

func (ArticleServiceServer) SetArticleVisibility(context.Context, *pb.SetArticleVisibilityRequest) (*pb.StatusResponse, error) {

}

func (ArticleServiceServer) ToggleArticleLike(context.Context, *pb.ToggleArticleLikeRequest) (*pb.StatusResponse, error) {

}
