package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	isPubilc := c.Locals("isPublic").(bool)

	if c.Method() == "GET" && isPubilc {
		return c.Next()
	}

	var accessToken string
	var err error
	if accessToken, err = getAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	// TODO : use websocket connection to authorize user
	log.Println(accessToken)
	username := "onee-only"
	author := c.Locals("author").(string)
	if username != author {
		return fiber.NewError(http.StatusForbidden, "author and requested user did not match")
	}

	return c.Next()
}

func getAccessFromHeader(header map[string]string) (string, error) {
	tokenRaw := header["Authorization"]
	tokenArr := strings.Split(tokenRaw, " ")
	if tokenArr[0] != "Bearer" || tokenArr[1] == "" {
		return "", fiber.NewError(http.StatusUnauthorized, "invalid auth method or token")
	}
	return tokenArr[1], nil
}
