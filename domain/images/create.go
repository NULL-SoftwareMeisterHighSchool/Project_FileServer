package images

import (
	"mime/multipart"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	image_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(userID, articleID uint, image *multipart.FileHeader) (string, *fiber.Error) {

	var err *fiber.Error

	name, extension := getNameAndExtension(image.Filename)
	if ok := checkExtension(extension, config.IMAGE_EXTENSIONS); !ok {
		return "", errors.ErrInvalidImageExtension
	}

	storage := storages.Get()

	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return "", errors.ErrNoPermission
	}

	var url string
	if url, err = storage.UploadImage(userID, name, extension, image); err != nil {
		return "", errors.CreateUnkownErr(err)
	}

	if err := image_repo.CreateImage(articleID, url); err != nil {
		return "", errors.CreateUnkownErr(err)
	}

	return url, nil
}
