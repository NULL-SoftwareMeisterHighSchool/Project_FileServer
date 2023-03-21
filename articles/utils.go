package articles

import "fmt"

func getArticlePath(username string, id int) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", username, id)
}
