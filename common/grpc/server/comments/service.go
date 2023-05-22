package comments_server

import (
	"context"

	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/comments"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
}

func (CommentServiceServer) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*empty.Empty, error) {
	if err := comments.CreateComment(
		uint(request.GetArticleID()),
		uint(request.GetAuthorID()),
		uint(request.GetReplyTo()),
		uint(request.GetMentionUserID()),
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
	comments, err := comments.GetCommentsFromArticle(uint(request.GetUserID()), uint(request.GetArticleID()))
	if err != nil {
		return nil, errors.StatusForError(err)
	}

	resElems := convertIntoCommentResElems(comments)

	return &pb.GetCommentsByArticleIDResponse{
		Comments: resElems,
	}, nil
}

func convertIntoCommentResElems(comments []*comment_repo.CommentWithReplyCount) []*pb.CommentElem {
	var resElems []*pb.CommentElem

	for _, comment := range comments {

		resElems = append(resElems, &pb.CommentElem{
			CommentID:  uint32(comment.ID),
			ReplyCount: uint32(comment.ReplyCount),
			AuthorID:   uint32(comment.AuthorID),
			CreatedAt:  timestamppb.New(comment.CreatedAt),
			Body:       comment.Body,
		})
	}

	return resElems
}

func (CommentServiceServer) GetRepliesByCommentID(context.Context, *pb.GetRepliesByCommentIDRequest) (*pb.GetRepliesByCommentIDResponse, error) {

}
