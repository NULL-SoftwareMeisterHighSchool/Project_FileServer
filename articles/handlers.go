package articles

import (
	"net/http"
	"os"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)

	err := c.SendFile(getArticlePath(author, id))
	if err != nil {
		return errors.NotFoundError
	}
	return nil
}

func CreateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)

	if articleExistsByPath(articlePath) {
		return errors.ConflictError
	}

	body := util.SanitizeXSS(c.Body())
	article := db.CreateArticle(id, body)
	db.Save(article)

	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)
	body := util.SanitizeXSS(c.Body())

	if !articleExistsByPath(articlePath) {
		return errors.NotFoundError
	}

	article := db.CreateArticle(id, body)

	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return err
	}

	// TODO : should update article entity, and images
	return c.Status(http.StatusOK).JSON()
}

func DeleteArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)

	if err := os.Remove(articlePath); err != nil {
		return errors.NotFoundError
	}
	// TODO : should delete article entity, and images

	return c.SendStatus(http.StatusNoContent)
}
