package grpc_client

import (
	core_dto "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/core/dto"
	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromToken(token string) uint {
	// TODO: should use grpc to retreive it
	return 1
}

func GetArticleInfoByID(id uint) (*core_dto.ArticleInfoRes, *fiber.Error) {
	// TODO: should use grpc to retreive it
	return nil, nil
}
