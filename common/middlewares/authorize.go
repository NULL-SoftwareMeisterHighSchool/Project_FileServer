package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/client/ws"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeMiddleware(c *fiber.Ctx) error {
	var err *fiber.Error
	isPubilc := c.Locals("isPublic").(bool)

	if c.Method() == "GET" && isPubilc {
		return c.Next()
	}

	var accessToken string
	if accessToken, err = util.ExtractAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	var userID uint
	if userID = ws.RequestUserIDFromToken(accessToken); userID == 0 {
		return errors.ErrAuthFailed
	}

	c.Locals("userID", userID)
	return c.Next()
}
