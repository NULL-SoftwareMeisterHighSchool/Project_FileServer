package github_client

import (
	"context"
	"log"
	"time"
)

const isoLayout = "2006-01-02T15:04:05-0700"

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
		"from":  from.UTC().Format(isoLayout),
		"to":    to.UTC().Format(isoLayout),
	}

	var query queryGithubStat
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	var (
		contCol     = query.User.ContributionsCollection
		starCnt     = query.User.Repositories.Nodes.Stargazers.TotalCount
		issueCnt    = query.User.Issues.TotalCount
		prCnt       = query.User.PullRequests.TotalCount
		contRepoCnt = query.User.PullRequests.TotalCount
	)

	return GithubInfo{
		ContributionCount:    uint(contCol.RestrictedContributionsCount + contCol.RestrictedContributionsCount),
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
