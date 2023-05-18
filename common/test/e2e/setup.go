package e2e_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func Setup(t *testing.T) {
	if err := godotenv.Load("../../.test.env"); err != nil {
		t.Fatalf("cannot load .test.env: %v", err)
	}
	config.LoadEnv()
	database.Connect()

	os.MkdirAll("./contents/images", 0755)

	t.Cleanup(func() {
		tearDown(t)
	})
}

func SetupgRPC(t *testing.T, register func(baseServer *grpc.Server)) (conn *grpc.ClientConn, close func()) {

	buffer := 1024 * 1024
	lis := bufconn.Listen(buffer)
	baseServer := grpc.NewServer()

	register(baseServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}

	close = func() {
		err := lis.Close()
		if err != nil {
			log.Fatalf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	return conn, close
}
