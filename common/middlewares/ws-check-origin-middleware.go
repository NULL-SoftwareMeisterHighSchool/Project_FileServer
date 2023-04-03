package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

func CheckOriginMiddleware(c *fiber.Ctx) error {
	if c.Get("host") != config.WS_REQUEST_ORIGIN {
		return errors.ErrInvalidOrigin
	}
	return c.Next()
}
