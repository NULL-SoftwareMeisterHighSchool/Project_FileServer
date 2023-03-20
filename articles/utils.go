package articles

import "fmt"

func getArticlePath(username string, id string) string {
	return fmt.Sprintf("./assets/%s/articles/%s.md", username, id)
}
