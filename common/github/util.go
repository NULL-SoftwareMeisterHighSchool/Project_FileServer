package github_client

import "time"

func getStarCountFrom(nodes []node) uint {
	var cnt uint = 0
	for _, node := range nodes {
		cnt += uint(node.Stargazers.TotalCount)
	}
	return cnt
}

func datetimeFrom(t time.Time) DateTime {
	return DateTime(t.UTC().Format(time.RFC3339))
}
