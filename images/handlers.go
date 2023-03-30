package images

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/storages"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/users"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	// should improve this duplication
	var accessToken string
	var err error
	if accessToken, err = util.ExtractAccessFromHeader(c.GetReqHeaders()); err != nil {
		return err
	}

	username := users.GetUsernameFromToken(accessToken)
	if username == "" {
		return errors.AuthFailedError
	}
	storage := storages.Get()
	
	
}
