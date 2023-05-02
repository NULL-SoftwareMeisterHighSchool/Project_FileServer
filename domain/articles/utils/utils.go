package article_utils

import (
	"strings"
)

func filterDeletableImageURLByEndpoint(urls []string, endpoint string) []string {
	shouldDelete := []string{}
	for _, url := range urls {
		if strings.HasPrefix(url, endpoint) {
			shouldDelete = append(shouldDelete, url)
		}
	}
	return shouldDelete
}
