package users_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/users"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserEventServiceServer
}

func (UserServiceServer) PublishUserCreated(ctx context.Context, event *pb.CreateUserEvent) (*empty.Empty, error) {

	if err := users.CreateUser(uint(event.GetUserID())); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (UserServiceServer) PublishUserDeleted(ctx context.Context, event *pb.DeleteUserEvent) (*empty.Empty, error) {

	if err := users.DeleteUser(uint(event.GetUserID())); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
