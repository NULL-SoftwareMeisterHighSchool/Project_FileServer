package github_client

import "time"

type node struct {
	Stargazers struct {
		TotalCount uint32
	}
}

type repositories struct {
	Nodes    []node
	PageInfo struct {
		EndCursor   string
		HasNextPage bool
	}
}

type contributionsCollection struct {
	TotalCommitContributions uint32
}

type queryGithubStat struct {
	User struct {
		ContributionsCollection   contributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
		RepositoriesContributedTo struct {
			TotalCount uint32
		}
		PullRequests struct {
			TotalCount uint32
		}
		Issues struct {
			TotalCount uint32
		}
		Repositories repositories `graphql:"repositories(first: 100, after: $after,orderBy: {direction:DESC, field: STARGAZERS}, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $login)"`
}

type queryGetUserJoinedAt struct {
	User struct {
		CreatedAt time.Time
	} `graphql:"user(login: $login)"`
}

type queryUserRepositoriesAndCommits struct {
	User struct {
		ContributionsCollection contributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
		Repositories            repositories            `graphql:"repositories(first: 100, after: $after,orderBy: {direction:DESC, field: STARGAZERS}, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $login)"`
}

type queryUserRepositories struct {
	User struct {
		Repositories repositories `graphql:"repositories(first: 100, after: $after,orderBy: {direction:DESC, field: STARGAZERS}, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $login)"`
}

type queryUserCommits struct {
	User struct {
		ContributionsCollection contributionsCollection `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $login)"`
}
