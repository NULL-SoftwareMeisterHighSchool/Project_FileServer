package images

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/users"
	"github.com/gofiber/fiber/v2"
)

func GetImage(c *fiber.Ctx) error {
	storage := storages.Get()
}

func UploadImage(c *fiber.Ctx) error {
	// should improve this duplication
	var accessToken string
	var err error
	if accessToken, err = util.ExtractAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	username := users.GetUsernameFromToken(accessToken)
	storage := storages.Get()
	
}
