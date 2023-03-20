package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetParamMiddleWare(c *fiber.Ctx) error {
	var id string
	if id = c.Params("id"); id == "" {
		return fiber.NewError(http.StatusNotFound, "id should be valid")
	}
	c.Locals("id", id)
	return c.Next()
}
