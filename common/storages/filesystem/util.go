package filesystem

import (
	"fmt"
	"os"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

func getImagePath(authorID uint, suffix string) string {
	return fmt.Sprintf("./contents/images/%d/%s", authorID, suffix)
}

func getArticlePath(authorID uint, id uint) string {
	return fmt.Sprintf("./contents/articles/%d/%d.md", authorID, id)
}

func articleExistsByPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getFSArticleURL(authorID uint, fullName string) string {
	return fmt.Sprintf("%s/%d/%s", config.IMAGE_HOST_ENDPOINT, authorID, fullName)
}
