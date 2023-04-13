package github_client

import (
	"context"
	"time"
)

const isoLayout = "2006-01-02T15:04:05-0700"

func GetUserContributionCount(login string, from, to *time.Time) int {
	client := getClient()

	variables := map[string]interface{}{
		"login": login,
		"from":  from.UTC().Format(isoLayout),
		"to":    to.UTC().Format(isoLayout),
	}

	var query queryCountUserContribution
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	return query.
		User.
		ContributionsCollection.
		ContributionCalendar.
		TotalContributions
}

func GetUserJoinedAt(login string) time.Time {
	client := getClient()

	variables := map[string]interface{}{
		"login": login,
	}

	var query queryGetUserJoinedAt
	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	return query.User.CreatedAt
}
