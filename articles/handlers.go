package articles

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/db"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)

	err := c.SendFile(getArticlePath(author, id))
	if err != nil {
		return fiber.NewError(
			http.StatusNotFound, fmt.Sprintf("article with id: %d, author: %s not found", id, author),
		)
	}
	return nil
}

func CreateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)

	if articleExistsByPath(articlePath) {
		return fiber.NewError(http.StatusConflict, "article already exists")
	}

	body := util.SanitizeXSS(c.Body())
	article := db.CreateArticleFromBody(body)

	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)

	if !articleExistsByPath(articlePath) {
		return fiber.NewError(http.StatusConflict, "article does not exist")
	}

	body := util.SanitizeXSS(c.Body())
	if err := os.WriteFile(articlePath, body, 0666); err != nil {
		return err
	}

	// TODO : should update article entity, and images
	return c.SendStatus(http.StatusAccepted)
}

func DeleteArticle(c *fiber.Ctx) error {
	id, author := getIdAndAuthor(c)
	articlePath := getArticlePath(author, id)

	if err := os.Remove(articlePath); err != nil {
		return fiber.NewError(http.StatusNotFound, "article does not exist")
	}
	// TODO : should delete article entity, and images

	return c.SendStatus(http.StatusNoContent)
}
