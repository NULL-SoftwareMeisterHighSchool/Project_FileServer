package users_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
)

type UserServiceServer struct {
	pb.UnimplementedUserEventServiceServer
}

func (UserServiceServer) PublishUserCreated(context.Context, *pb.CreateUserEvent) (*pb.Nothing, error) {

}

func (UserServiceServer) PublishUserDeleted(context.Context, *pb.DeleteUserEvent) (*pb.Nothing, error) {

}
