package images

import (
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func getNameAndExtension(filename string) (name, extension string) {
	splitted := strings.Split(filename, ".")
	extension = splitted[len(splitted)-1]
	name = util.CreateUUID()
	return
}

func checkExtension(candidate string, allowedExtensions []string) bool {
	for _, extension := range allowedExtensions {
		if candidate == extension {
			return true
		}
	}
	return false
}
