package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	accessToken, err := getAccessFromHeader(c.GetReqHeaders())
	if err != nil && c.Method() != "GET" {
		return err
	}

	// TODO : use websocket connection to authorize user
	log.Println(accessToken)
	c.Locals("author", "onee-only")

	return nil
}

func getAccessFromHeader(header map[string]string) (string, error) {
	tokenRaw := header["Authorization"]
	tokenArr := strings.Split(tokenRaw, " ")
	if tokenArr[0] != "Bearer" || tokenArr[1] == "" {
		return "", fiber.NewError(http.StatusUnauthorized, "invalid auth method or token")
	}
	return tokenArr[1], nil
}
