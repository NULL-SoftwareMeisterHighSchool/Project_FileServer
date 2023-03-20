package common

import "github.com/gofiber/fiber/v2"

func GetParamMiddleWare(c *fiber.Ctx) error {
	var id string
	if err := getParam(&id, "id", c); err != nil {
		return err
	}
	c.Locals("id", id)
	return nil
}
