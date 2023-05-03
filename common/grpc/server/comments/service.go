package comments_server

import (
	"context"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
}

func (CommentServiceServer) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*empty.Empty, error) {
	if err := comments.CreateComment(
		uint(request.GetArticleID()),
		uint(request.GetAuthorID()),
		uint(request.GetReplyTo()),
		request.GetBody(),
	); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (CommentServiceServer) DeleteCommnet(ctx context.Context, request *pb.DeleteCommnetRequest) (*empty.Empty, error) {
	if err := comments.DeleteComment(
		uint(request.GetCommentID()),
		uint(request.GetUserID()),
		uint(request.GetArticleID()),
	); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &emptypb.Empty{}, nil
}

func (CommentServiceServer) GetCommentsByArticleID(ctx context.Context, request *pb.GetCommentsByArticleIDRequest) (*pb.GetCommentsByArticleIDResponse, error) {

}
