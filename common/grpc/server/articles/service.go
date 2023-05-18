package articles_server

import (
	"context"

	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
}

func (ArticleServiceServer) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	res, err := articles.CreateArticle(uint(request.GetAuthorID()), request.GetTitle(), request.GetType())
	if err != nil {
		return nil, errors.StatusForError(err)
	}
	return res, nil
}

func (ArticleServiceServer) ListArticle(ctx context.Context, request *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {
	res, err := articles.ListArticles(
		uint(request.GetOffset()),
		uint(request.GetAmount()),
		article_repo.ArticleOrder(request.GetOrder()),
		article_repo.ArticleTypeOption(request.GetType()),
		uint(request.GetAuthorID()),
		request.GetDuration(),
		request.GetQuery(),
	)

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return res, nil
}

func (ArticleServiceServer) ListArticleByAuthor(ctx context.Context, request *pb.ListArticleByAuthorRequest) (*pb.ListArticleResponse, error) {
	res, err := articles.ListArticlesByAuthor(
		uint(request.GetOffset()),
		uint(request.GetAmount()),
		article_repo.ArticleOrder(request.GetOrder()),
		article_repo.ArticleTypeOption(request.GetType()),
		uint(request.GetAuthorID()),
		uint(request.GetUserID()),
		request.GetDuration(),
		request.GetQuery(),
		request.IsPrivate,
	)

	if err != nil {
		return nil, errors.StatusForError(err)
	}

	return res, nil
}

func (ArticleServiceServer) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	res, err := articles.GetArticle(uint(request.GetArticleID()), uint(request.GetUserID()))

	if err != nil {
		return nil, errors.StatusForError(err)
	}
	return res, nil
}

func (ArticleServiceServer) UpdateArticleBody(ctx context.Context, request *pb.UpdateArticleBodyRequest) (*empty.Empty, error) {
	if err := articles.UpdateArticleBody(uint(request.GetArticleID()), uint(request.GetUserID()), []byte(request.GetBody())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (ArticleServiceServer) UpdateArticleTitle(ctx context.Context, request *pb.UpdateArticleTitleRequest) (*empty.Empty, error) {
	if err := articles.UpdateArticleTitle(uint(request.GetArticleID()), uint(request.GetUserID()), string(request.GetTitle())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (ArticleServiceServer) DeleteArticle(ctx context.Context, request *pb.DeleteArticleRequest) (*empty.Empty, error) {
	if err := articles.DeleteByID(uint(request.GetArticleID()), uint(request.GetUserID())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (ArticleServiceServer) SetArticleVisibility(ctx context.Context, request *pb.SetArticleVisibilityRequest) (*empty.Empty, error) {
	if err := articles.UpdateArticleVisibility(uint(request.GetUserID()), uint(request.GetArticleID()), request.GetIsPrivate()); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (ArticleServiceServer) ToggleArticleLike(ctx context.Context, request *pb.ToggleArticleLikeRequest) (*empty.Empty, error) {
	if err := articles.DeleteByID(uint(request.GetArticleID()), uint(request.GetUserID())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}
