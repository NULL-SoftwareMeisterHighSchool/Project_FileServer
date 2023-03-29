package main

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/middlewares"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/images"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()
	db.Connect()

	app := initApp()
	app.Listen(":8080")
}

func initApp() *fiber.App {
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

		// images
		imagesRouter := app.Group("images")
		{
			imagesRouter.Get("/:name", images.GetImage)
			imagesRouter.Post("", images.UploadImage)
		}

	}
	return app
}
