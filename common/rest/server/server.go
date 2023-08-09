package rest_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/middlewares"
	image_controller "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest/server/images"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Listen() {

	appConfig := fiber.Config{
		ErrorHandler: errors.CustomErrorHandler,
	}

	app := fiber.New(appConfig)
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	{
		app.Get("ping", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON(fiber.Map{"message": "pong"})
		})

		imagesRouter := app.Group("images")
		{
			imagesRouter.Post("",
				middlewares.AuthorizeMiddleware,
				image_controller.UploadImage,
			)
			// filesystem setting. can be deleted if it uses s3 or something.
			// imagesRouter.Static("", "./contents/images")
		}

	}
	if err := app.Listen(fmt.Sprintf(":%s", config.REST_PORT)); err != nil {
		log.Fatal(err)
	}
}
