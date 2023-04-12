package middlewares

import (
	core_client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/client/core"
	"github.com/gofiber/fiber/v2"
)

func GetArticleInfoMiddleware(c *fiber.Ctx) error {
	articleInfo, err := core_client.Get().GetArticleInfoByID(c.Locals("id").(uint))
	if err != nil {
		return err
	}
	c.Locals("authorID", articleInfo.AuthorID)
	c.Locals("isPublic", articleInfo.IsPublic)
	return c.Next()
}
