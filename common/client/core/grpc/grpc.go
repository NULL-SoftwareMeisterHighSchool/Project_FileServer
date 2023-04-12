package grpc_client

import (
	core_dto "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/client/core/dto"
	"github.com/gofiber/fiber/v2"
)

type gc struct{}

func Get() *gc {
	return &gc{}
}

func (gc) GetUserIDFromToken(token string) uint {
	// TODO: should use websocket connection to retreive it
	return 1
}

func (gc) GetArticleInfoByID(id uint) (*core_dto.ArticleInfo, *fiber.Error) {
	// TODO: should use websocket connection to retreive it
	return nil, nil
}
