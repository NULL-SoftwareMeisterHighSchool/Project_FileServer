package github_client

import "time"

type Node struct {
	Stargazers struct {
		TotalCount uint32
	}
}

type queryGithubStat struct {
	User struct {
		ContributionsCollection struct {
			RestrictedContributionsCount uint32
			TotalCommitContributions     uint32
		} `graphql:"contributionsCollection(from: $from, to: $to)"`
		RepositoriesContributedTo struct {
			TotalCount uint32
		}
		PullRequests struct {
			TotalCount uint32
		}
		Issues struct {
			TotalCount uint32
		}
		Repositories struct {
			Nodes []Node
		} `graphql:"repositories(first: 10000, ownerAffiliations: OWNER"`
	} `graphql:"user(login: $login)"`
}

type queryGetUserJoinedAt struct {
	User struct {
		CreatedAt time.Time
	} `graphql:"user(login: $login)"`
}
