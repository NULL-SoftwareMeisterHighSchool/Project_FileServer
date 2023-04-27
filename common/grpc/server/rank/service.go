package rank_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/rank"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/rank"
)

type RankServiceServer struct {
	pb.UnimplementedRankServiceServer
}

func (RankServiceServer) GetContributionCount(ctx context.Context, request *pb.GetContributionCountRequest) (*pb.GetCountributionCountResponse, error) {
	return rank.GetUsersContribution(request.GetUsers())
}
