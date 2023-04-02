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
		articlesRouter := app.Group("articles")
		articlesRouter.Get("", articles.GetArticleInfo)

		articleRouter := articlesRouter.Group(":id")
		{
			articleRouter.Use(middlewares.GetParamMiddleware)
			articleRouter.Use(middlewares.GetArticleInfoMiddleware)
			articleRouter.Use(middlewares.AuthMiddleware)

			articleRouter.Get("", articles.GetArticle)
			articleRouter.Post("", articles.CreateArticle)
			articleRouter.Put("", articles.UpdateArticle)
			articleRouter.Delete("", articles.DeleteArticle)
		}

		// images
		app.Post("images", images.UploadImage)
		// filesystem setting. can be deleted if it uses s3 or something.
		app.Static("images", "./contents/images")

	}
	return app
}
