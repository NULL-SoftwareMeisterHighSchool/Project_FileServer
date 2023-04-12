package core_client

import (
	core_dto "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/client/core/dto"
	grpc_client "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/client/core/grpc"
	"github.com/gofiber/fiber/v2"
)

type CoreService interface {
	GetUserIDFromToken(token string) uint
	GetArticleInfoByID(id uint) (*core_dto.ArticleInfo, *fiber.Error)
}

var coreService CoreService = grpc_client.Get()

func Get() CoreService {
	return coreService
}
