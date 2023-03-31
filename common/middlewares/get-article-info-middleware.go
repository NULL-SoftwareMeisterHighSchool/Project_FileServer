package middlewares

import "github.com/gofiber/fiber/v2"

func GetArticleInfoMiddleware(c *fiber.Ctx) error {
	// TODO : get article info
	var authorID uint = 1
	isPublic := false
	c.Locals("authorID", authorID)
	c.Locals("isPublic", isPublic)
	return c.Next()
}
