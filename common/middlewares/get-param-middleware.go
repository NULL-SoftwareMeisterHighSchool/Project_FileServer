package middlewares

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/gofiber/fiber/v2"
)

func GetParamMiddleware(c *fiber.Ctx) error {
	var id int
	var err error
	if id, err = c.ParamsInt("id"); err != nil {
		return errors.InvalidIDError
	}
	c.Locals("id", id)
	return c.Next()
}
