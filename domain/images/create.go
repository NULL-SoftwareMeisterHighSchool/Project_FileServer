package images

import (
	"mime/multipart"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	image_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(userID uint, image *multipart.FileHeader) (string, *fiber.Error) {

	var err *fiber.Error

	name, extension := getNameAndExtension(image.Filename)
	if ok := checkExtension(extension, config.IMAGE_EXTENSIONS); !ok {
		return "", errors.ErrInvalidImageExtension
	}

	storage := storages.Get()

	var url string
	if url, err = storage.UploadImage(name, extension, image); err != nil {
		return "", errors.CreateUnkownErr(err)
	}

	if err := image_repo.CreateImage(url); err != nil {
		return "", errors.CreateUnkownErr(err)
	}

	return url, nil
}
