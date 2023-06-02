package users_test

import (
	"context"
	"log"
	"testing"

	users_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
	users_server "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/users"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
	"google.golang.org/grpc"
)

func userBeforeEach(t *testing.T) users_pb.UserEventServiceClient {
	conn, close := e2e_test.SetupgRPC(t, func(baseServer *grpc.Server) {
		users_pb.RegisterUserEventServiceServer(baseServer, users_server.UserServiceServer{})
	})

	client := users_pb.NewUserEventServiceClient(conn)
	t.Cleanup(close)
	return client
}

func TestGetUsersGithubStats(t *testing.T) {
	e2e_test.Setup(t)
	client := userBeforeEach(t)

	res, _ := client.GetGithubStats(context.Background(), &users_pb.GetGithubStatsRequest{
		Users: []*users_pb.UserInfo{{
			UserID:    uint32(1),
			UserLogin: "onee-only",
		}},
	})

	log.Println(res)
}
