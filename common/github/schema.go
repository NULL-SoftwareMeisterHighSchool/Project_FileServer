package github_client

import "time"

type queryCountUserContribution struct {
	User struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions int
			} `graphql:"contributionCalendar(from: $from, to: $to)"`
		}
	} `graphql:"user(login: $login)"`
}

type queryGetUserJoinedAt struct {
	User struct {
		CreatedAt time.Time
	} `graphql:"user(login: $login)"`
}
