package comments_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
	"github.com/golang/protobuf/ptypes/empty"
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
}

func (CommentServiceServer) CreateComment(context.Context, *pb.CreateCommentRequest) (*empty.Empty, error) {

}

func (CommentServiceServer) DeleteCommnet(context.Context, *pb.DeleteCommnetRequest) (*empty.Empty, error) {

}

func (CommentServiceServer) GetCommentsByArticleID(context.Context, *pb.GetCommentsByArticleIDRequest) (*pb.GetCommentsByArticleIDResponse, error) {

}
