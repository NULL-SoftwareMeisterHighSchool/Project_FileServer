package middlewares

import (
	"strconv"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

func GetParamMiddleware(c *fiber.Ctx) error {
	var id uint64
	var err error
	if id, err = strconv.ParseUint(c.Params("id"), 10, 32); err != nil {
		return errors.ErrInvalidID
	}
	c.Locals("id", uint(id))
	return c.Next()
}
