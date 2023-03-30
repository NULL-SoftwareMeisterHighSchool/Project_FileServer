package images

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

func getNameAndExtension(filename string) (name, extension string) {
	splitted := strings.Split(filename, ".")
	extension = splitted[len(splitted)-1]
	name = createName()
	return
}

func createName() string {
	randHex := fmt.Sprintf("%x", rand.Uint64())
	var runes []rune
	for i, char := range randHex {
		if i%4 == 0 && i != 0 {
			runes = append(runes, '-')
		}
		runes = append(runes, char)
	}
	return string(runes)
}

func checkExtension(candidate string) bool {
	for _, extension := range config.IMAGE_EXTENSIONS {
		if candidate == extension {
			return true
		}
	}
	return false
}
