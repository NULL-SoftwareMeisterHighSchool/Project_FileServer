package comments_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
}

func (CommentServiceServer) CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.StatusResponse, error) {

}
func (CommentServiceServer) DeleteCommnet(context.Context, *pb.DeleteCommnetRequest) (*pb.StatusResponse, error) {

}
func (CommentServiceServer) GetCommentsByArticleID(context.Context, *pb.GetCommentsByArticleIDRequest) (*pb.GetCommentsByArticleIDResponse, error) {

}
