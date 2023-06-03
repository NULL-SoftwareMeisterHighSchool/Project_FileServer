package github_client

import (
	"context"
	"log"
	"time"

	"github.com/hasura/go-graphql-client"
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

type pageinfo struct {
	hasNext   bool
	endCursor string
}

func GetUserContributionCount(login string, from, end time.Time) GithubInfo {

	var (
		githubInfo *GithubInfo
		pageInfo   *pageinfo

		next = from.AddDate(1, 0, 0)
	)
	client := getClient()

	variables := map[string]interface{}{
		"login": login,
		"from":  datetimeFrom(from),
		"to":    datetimeFrom(next),
		"after": nil,
	}

	githubInfo, pageInfo = queryAllGithubStats(client, variables)

loop:
	for {
		var (
			commitCnt, starCnt uint = 0, 0

			hasNextPage    = pageInfo.hasNext
			nextIsBeforeTo = next.Before(end)
		)

		from = next
		next = next.AddDate(1, 0, 0)

		variables["after"] = pageInfo.endCursor
		variables["from"] = datetimeFrom(from)
		variables["to"] = datetimeFrom(next)

		switch {
		case hasNextPage && nextIsBeforeTo:
			// query commit & star
			commitCnt, starCnt, pageInfo = queryNextAll(client, variables)
		case hasNextPage:
			// query star
			starCnt, pageInfo = queryNextStarCnt(client, variables)
		case nextIsBeforeTo:
			// query commit
			commitCnt = queryNextCommitCnt(client, variables)
		default:
			break loop
		}

		githubInfo.ContributionCount += commitCnt
		githubInfo.StarCount += starCnt
	}

	return *githubInfo
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

func queryAllGithubStats(client *graphql.Client, variables map[string]interface{}) (*GithubInfo, *pageinfo) {

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

	githubInfo := &GithubInfo{
		ContributionCount:    uint(commitCnt),
		StarCount:            uint(starCnt),
		IssueCount:           uint(issueCnt),
		PullRequestCount:     uint(prCnt),
		ContributedRepoCount: uint(contRepoCnt),
	}

	pageInfo := &pageinfo{
		hasNext:   query.User.Repositories.PageInfo.HasNextPage,
		endCursor: query.User.Repositories.PageInfo.EndCursor,
	}

	return githubInfo, pageInfo
}

func queryNextAll(client *graphql.Client, variables map[string]interface{}) (commitCnt, starCnt uint, pageInfo *pageinfo) {
	var query queryUserRepositoriesAndCommits
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	commitCnt = uint(query.User.ContributionsCollection.TotalCommitContributions)
	starCnt = getStarCountFrom(query.User.Repositories.Nodes)
	pageInfo = &pageinfo{
		hasNext:   query.User.Repositories.PageInfo.HasNextPage,
		endCursor: query.User.Repositories.PageInfo.EndCursor,
	}

	return commitCnt, starCnt, pageInfo
}

func queryNextStarCnt(client *graphql.Client, variables map[string]interface{}) (starCnt uint, pageInfo *pageinfo) {
	var query queryUserRepositories
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	starCnt = getStarCountFrom(query.User.Repositories.Nodes)
	pageInfo = &pageinfo{
		hasNext:   query.User.Repositories.PageInfo.HasNextPage,
		endCursor: query.User.Repositories.PageInfo.EndCursor,
	}

	return starCnt, pageInfo
}

func queryNextCommitCnt(client *graphql.Client, variables map[string]interface{}) (commitCnt uint) {
	var query queryUserCommits
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	commitCnt = uint(query.User.ContributionsCollection.TotalCommitContributions)

	return commitCnt
}
