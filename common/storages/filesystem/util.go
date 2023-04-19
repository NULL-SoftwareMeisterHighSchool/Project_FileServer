package filesystem

import (
	"fmt"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

func getImagePath(authorID uint, suffix string) string {
	return fmt.Sprintf("./contents/images/%d/%s", authorID, suffix)
}

func getImageURL(authorID uint, fullName string) string {
	return fmt.Sprintf("%s/%d/%s", config.IMAGE_HOST_ENDPOINT, authorID, fullName)
}
