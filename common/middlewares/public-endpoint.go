package middlewares

import "github.com/gofiber/fiber/v2"

func SetPublicTrue(c *fiber.Ctx) error {
	c.Locals("isPublic", true)
	return c.Next()
}
