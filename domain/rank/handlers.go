package rank

import (
	"time"

	github_client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/github"
)

func GetUserContribution(login string) int {
	joinedAt := github_client.GetUserJoinedAt(login)
	now := time.Now()
	return github_client.GetUserContributionCount(login, &joinedAt, &now)
}
