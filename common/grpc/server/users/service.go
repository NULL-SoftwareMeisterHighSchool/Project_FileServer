package users_server

import (
	"context"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/users"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserServiceServer struct {
	pb.UnimplementedUserEventServiceServer
}

func (UserServiceServer) PublishUserCreated(ctx context.Context, event *pb.CreateUserEvent) (*empty.Empty, error) {

	if err := users.CreateUser(uint(event.GetUserID())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &empty.Empty{}, nil
}

func (UserServiceServer) PublishUserDeleted(ctx context.Context, event *pb.DeleteUserEvent) (*empty.Empty, error) {

	if err := users.DeleteUser(uint(event.GetUserID())); err != nil {
		return nil, errors.StatusForError(err)
	}

	return &empty.Empty{}, nil
}

func (UserServiceServer) GetGithubStats(ctx context.Context, request *pb.GetGithubStatsRequest) (*pb.GetGithubStatsResponse, error) {

	stats, err := users.GetUsersGithubStats(request.GetUsers())
	if err != nil {
		return nil, errors.StatusForError(err)
	}

	return stats, nil
}
