package rest

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/middlewares"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/images"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitApp() *fiber.App {

	// config
	config := fiber.Config{
		ErrorHandler: errors.CustomErrorHandler,
	}

	app := fiber.New(config)
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
			articleRouter.Use(middlewares.AuthorizeMiddleware)
			articleRouter.Use(middlewares.AuthenticateMiddleware)

			articleRouter.Get("", articles.GetArticle)
			articleRouter.Post("", articles.CreateArticle)
			articleRouter.Put("", articles.UpdateArticle)
			articleRouter.Delete("", articles.DeleteArticle)
		}

		// images
		imagesRouter := app.Group("images")
		{
			imagesRouter.Post("",
				middlewares.SetPublicTrue,
				middlewares.AuthorizeMiddleware,
				images.UploadImage,
			)
			// filesystem setting. can be deleted if it uses s3 or something.
			imagesRouter.Static("", "./contents/images")
		}

	}
	return app
}
