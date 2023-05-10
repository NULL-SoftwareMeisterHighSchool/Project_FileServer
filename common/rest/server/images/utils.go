package image_controller

import (
	"mime/multipart"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

func getImageFromFormFile(c *fiber.Ctx) (*multipart.FileHeader, *fiber.Error) {
	image, err := c.FormFile("image")
	if err != nil {
		return nil, errors.ErrImageNotReceived
	}
	return image, nil
}
