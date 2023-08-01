package image_controller

import (
	"mime/multipart"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/images"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	var err *fiber.Error

	userID := c.Locals("userID").(uint)

	var image *multipart.FileHeader
	if image, err = getImageFromFormFile(c); err != nil {
		return err
	}

	url, err := images.UploadImage(userID, image)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"url": url})
}

func parseQuery(c *fiber.Ctx) (uint, *fiber.Error) {
	articleID := uint(c.QueryInt("articleID"))
	if articleID == 0 {
		return 0, errors.ErrInvalidArticleID
	}
	return articleID, nil
}
