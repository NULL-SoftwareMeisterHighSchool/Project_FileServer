package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/users"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	isPubilc := c.Locals("isPublic").(bool)

	if c.Method() == "GET" && isPubilc {
		return c.Next()
	}

	var accessToken string
	var err error
	if accessToken, err = util.ExtractAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	username := users.GetUsernameFromToken(accessToken)
	if username == "" {
		return errors.AuthFailedError
	}

	author := c.Locals("author").(string)
	if username != author {
		return errors.NoPermissionError
	}

	return c.Next()
}
