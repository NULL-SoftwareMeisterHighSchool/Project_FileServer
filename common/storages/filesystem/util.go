package filesystem

import (
	"fmt"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

func getImagePath(suffix string) string {
	return fmt.Sprintf("./contents/images/%s", suffix)
}

func getImageURL(fullName string) string {
	return fmt.Sprintf("%s/%s", config.IMAGE_HOST_ENDPOINT, fullName)
}
