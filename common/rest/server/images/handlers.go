package image_controller

import (
	"mime/multipart"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/images"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	var err *fiber.Error
	userID := c.Locals("userID").(uint)
	articleID := c.Locals("articleID").(uint)

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
