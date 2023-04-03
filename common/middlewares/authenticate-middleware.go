package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {

	userID := c.Locals("userID").(uint)
	authorID := c.Locals("authorID").(uint)

	if userID != authorID {
		return errors.ErrNoPermission
	}

	return c.Next()
}
