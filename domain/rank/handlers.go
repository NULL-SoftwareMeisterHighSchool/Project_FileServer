package rank

import (
	"time"

	github_client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/github"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/rank"
)

func requestContribution(login string, c chan<- uint32) {
	joinedAt := github_client.GetUserJoinedAt(login)
	now := time.Now()
	c <- github_client.GetUserContributionCount(login, &joinedAt, &now)
}

func GetUsersContribution(userInfoList []*pb.UserInfo) (*pb.GetCountributionCountResponse, error) {
	contMap := make(map[uint32]chan uint32)

	for _, userInfo := range userInfoList {
		c := make(chan uint32)
		contMap[userInfo.GetUserID()] = c
		go requestContribution(userInfo.GetUserLogin(), c)
	}

	var contElems []*pb.ContributionElement

	for userID, contInfoChan := range contMap {
		cnt := <-contInfoChan
		contElems = append(contElems, &pb.ContributionElement{
			UserID: userID,
			Count:  cnt,
		})
	}

	return &pb.GetCountributionCountResponse{
		ContElems: contElems,
	}, nil
}
