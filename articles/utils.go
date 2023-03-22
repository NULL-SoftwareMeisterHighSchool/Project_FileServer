package articles

import "fmt"

func getArticlePath(author string, id int) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}
