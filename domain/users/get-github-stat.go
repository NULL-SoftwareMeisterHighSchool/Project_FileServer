package users

import (
	"time"

	client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/github"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/users"
)

func requestGithubStat(userInfo *pb.UserInfo, c chan<- client.GithubInfo) {
	joinedAt := client.GetUserJoinedAt(userInfo.GetUserLogin())
	now := time.Now()

	stat := client.GetUserContributionCount(userInfo.GetUserLogin(), &joinedAt, &now)
	stat.UserID = uint(userInfo.GetUserID())
	c <- stat
}

func getGithubStatStream(userInfoList []*pb.UserInfo) chan client.GithubInfo {
	statStream := make(chan client.GithubInfo)

	for _, userInfo := range userInfoList {
		go requestGithubStat(userInfo, statStream)
	}

	return statStream
}

func GetUsersGithubStats(userInfoList []*pb.UserInfo) (*pb.GetGithubStatsResponse, error) {

	var statElems []*pb.GithubStats
	statStream := getGithubStatStream(userInfoList)

	for ghStat := range statStream {
		statElems = append(statElems, &pb.GithubStats{
			UserID:                     uint32(ghStat.UserID),
			ContributionCount:          uint32(ghStat.ContributionCount),
			StarCount:                  uint32(ghStat.StarCount),
			IssueCount:                 uint32(ghStat.IssueCount),
			PullRequestCount:           uint32(ghStat.PullRequestCount),
			ContributedRepositoryCount: uint32(ghStat.ContributedRepoCount),
		})
	}

	return &pb.GetGithubStatsResponse{
		StatElems: statElems,
	}, nil
}
