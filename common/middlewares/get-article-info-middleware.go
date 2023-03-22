package middlewares

import "github.com/gofiber/fiber/v2"

func GetArticleInfoMiddleware(c *fiber.Ctx) error {
	// TODO : get article info
	author := ""
	isPublic := false
	c.Locals("author", author)
	c.Locals("isPublic", isPublic)
	return c.Next()
}
