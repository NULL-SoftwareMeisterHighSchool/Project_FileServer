package articles

import (
	entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
)

func CreateArticle(authorID uint, title string, articleType pb.ArticleType) error {
	return article_repo.CreateArticle(authorID, title, entity.ArticleType(articleType))
}
