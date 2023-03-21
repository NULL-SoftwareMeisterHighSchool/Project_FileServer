package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetParamMiddleware(c *fiber.Ctx) error {
	var id int
	var err error
	if id, err = c.ParamsInt("id"); err != nil {
		return fiber.NewError(http.StatusBadRequest, "id should be valid")
	}
	c.Locals("id", id)
	return c.Next()
}
