package images

import (
	"mime/multipart"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/users"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	var err *fiber.Error

	var userID uint
	if userID, err := users.GetUserIDFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	storage := storages.Get()

	var image *multipart.FileHeader
	if image, err = getImageFromFormFile(c); err != nil {
		return err
	}

	name, extension := getNameAndExtension(image.Filename)
	if ok := checkExtension(extension); !ok {
		return errors.InvalidImageExtensionError
	}

	var url string
	if url, err = storage.UploadImage(username, name, extension, image); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"url": url})
}
