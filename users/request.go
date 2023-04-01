package users

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	"github.com/gofiber/fiber/v2"
)

func getUserIDFromToken(token string) uint {
	// TODO: should use websocket connection to retreive it
	return 1
}

func GetUserIDFromHeader(header map[string]string) (uint, *fiber.Error) {
	var accessToken string
	var err *fiber.Error
	if accessToken, err = util.ExtractAccessFromHeader(header); err != nil {
		return 0, err
	}
	userID := getUserIDFromToken(accessToken)
	if userID == 0 {
		return 0, errors.ErrAuthFailed
	}
	return userID, nil
}
