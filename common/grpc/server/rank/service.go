package rank_server

import (
	"context"

	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/rank"
)

type RankServiceServer struct {
	pb.UnimplementedRankServiceServer
}

func (RankServiceServer) GetContributionCount(context.Context, *pb.GetContributionCountRequest) (*pb.GetCountributionCountResponse, error) {
	
}
