package core_client

import "github.com/gofiber/fiber/v2"

type ArticleInfo struct {
	IsPublic bool
	AuthorID uint
}

func RequestUserIDFromToken(token string) uint {
	// TODO: should use websocket connection to retreive it
	return 1
}

func GetArticleInfoByID(id uint) (*ArticleInfo, *fiber.Error) {
	// TODO: should use websocket connection to retreive it
	return nil, nil
}
