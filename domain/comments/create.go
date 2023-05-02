package comments

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles"
)

func CreateComment(articleID, userID uint, replyTo uint, body string) error {
	if err := articles.CheckPrivateAndExists(userID, articleID); err != nil {
		return err
	}
	return comment_repo.CreateComment(articleID, userID, replyTo, body)
}
