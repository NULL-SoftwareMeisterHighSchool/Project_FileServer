package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	core_client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest/client/core"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeMiddleware(c *fiber.Ctx) error {
	var err *fiber.Error

	var accessToken string
	if accessToken, err = util.ExtractAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	var userID uint
	if userID = core_client.GetUserIDFromToken(accessToken); userID == 0 {
		return errors.ErrAuthFailed
	}

	c.Locals("userID", userID)
	return c.Next()
}
