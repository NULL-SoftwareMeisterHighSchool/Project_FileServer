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
	var articleID uint

	userID := c.Locals("userID").(uint)
	if articleID, err = parseQuery(c); err != nil {
		return err
	}

	var image *multipart.FileHeader
	if image, err = getImageFromFormFile(c); err != nil {
		return err
	}

	url, err := images.UploadImage(userID, articleID, image)
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
