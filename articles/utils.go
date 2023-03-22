package articles

import (
	"fmt"
	"os"
)

func getArticlePath(author string, id int) string {
	return fmt.Sprintf("./contents/%s/articles/%d.md", author, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
