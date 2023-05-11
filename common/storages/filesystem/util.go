package filesystem

import (
	"fmt"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

func getImagePath(suffix string) string {
	return fmt.Sprintf("./contents/images/%s", suffix)
}

func getImageURL(fullName string) string {
	return fmt.Sprintf("%s/%s", config.IMAGE_HOST_ENDPOINT, fullName)
}

func getSuffixFromURL(url string) string {
	suffix, _ := strings.CutPrefix(url, fmt.Sprintf("%s/", config.IMAGE_HOST_ENDPOINT))
	return suffix
}
