package main

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// routes
	{

		// ping pong
		app.Get("ping", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON(fiber.Map{"message": "pong"})
		})

		// articles
		articlesRouter := app.Group("articles/:id")
		articlesRouter.Use(middlewares.GetParamMiddleware)
		articlesRouter.Use(middlewares.GetArticleInfoMiddleware)
		articlesRouter.Use(middlewares.AuthMiddleware)
		{
			articlesRouter.Get("", articles.GetArticle)
			articlesRouter.Post("", articles.CreateArticle)
			articlesRouter.Put("", articles.UpdateArticle)
			articlesRouter.Delete("", articles.DeleteArticle)
		}

	}

	app.Listen(":8080")
}
