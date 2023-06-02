package github_client

import (
	"context"
	"log"
	"time"
)

type DateTime string

type GithubInfo struct {
	UserID               uint
	ContributionCount    uint
	StarCount            uint
	IssueCount           uint
	PullRequestCount     uint
	ContributedRepoCount uint
}

func GetUserContributionCount(login string, from, to *time.Time) GithubInfo {
	client := getClient()

	variables := map[string]interface{}{
		"login": login,
		"from":  DateTime(from.UTC().Format(time.RFC3339)),
		"to":    DateTime(to.UTC().Format(time.RFC3339)),
	}

	var query queryGithubStat
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	var (
		commitCnt   = query.User.ContributionsCollection.TotalCommitContributions
		starCnt     = getStarCountFrom(query.User.Repositories.Nodes)
		issueCnt    = query.User.Issues.TotalCount
		prCnt       = query.User.PullRequests.TotalCount
		contRepoCnt = query.User.PullRequests.TotalCount
	)

	return GithubInfo{
		ContributionCount:    uint(commitCnt),
		StarCount:            uint(starCnt),
		IssueCount:           uint(issueCnt),
		PullRequestCount:     uint(prCnt),
		ContributedRepoCount: uint(contRepoCnt),
	}
}

func GetUserJoinedAt(login string) time.Time {
	client := getClient()

	variables := map[string]interface{}{
		"login": login,
	}

	var query queryGetUserJoinedAt
	if err := client.Query(context.Background(), &query, variables); err != nil {
		log.Fatalf("failed to send query: %s", err)
	}

	return query.User.CreatedAt
}

func getStarCountFrom(nodes []Node) uint {
	var cnt uint = 0
	for _, node := range nodes {
		cnt += uint(node.Stargazers.TotalCount)
	}
	return cnt
}
