package main

import (
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/articles"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/middlewares"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/images"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func main() {
	config.LoadEnv()
	db.Connect()

	app := initApp()
	app.Listen(":8080")
}

func initApp() *fiber.App {

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

		// websocket
		app.Get("ws", middlewares.CheckOriginMiddleware, websocket.New(ws.Connect))

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
