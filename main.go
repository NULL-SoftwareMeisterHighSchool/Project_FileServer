package main

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// server
	app := fiber.New()
	{
		// route

		// ping pong
		app.Get("ping", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON(fiber.Map{"message": "pong"})
		})

		// articles
		articlesRouter := app.Group("articles")
		articlesRouter.Use(common.GetParamMiddleWare)
		// TODO : create auth middleware
		{
			articlesRouter.Get(":id", articles.GetArticle)
			articlesRouter.Post(":id", articles.CreateArticle)
			articlesRouter.Put(":id", articles.UpdateArticle)
			articlesRouter.Delete(":id", articles.DeleteArtcile)
		}

		app.Listen(":8080")
	}
}
