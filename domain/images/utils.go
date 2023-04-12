package images

import (
	"mime/multipart"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
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

func getImageFromFormFile(c *fiber.Ctx) (*multipart.FileHeader, *fiber.Error) {
	image, err := c.FormFile("image")
	if err != nil {
		return nil, errors.ErrImageNotReceived
	}
	return image, nil
}
