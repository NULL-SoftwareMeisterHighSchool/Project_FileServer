package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/users"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	isPubilc := c.Locals("isPublic").(bool)

	if c.Method() == "GET" && isPubilc {
		return c.Next()
	}

	var userID uint
	var err *fiber.Error
	if userID, err = users.GetUserIDFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	authorID := c.Locals("authorID").(uint)
	if userID != authorID {
		return errors.ErrNoPermission
	}

	return c.Next()
}
