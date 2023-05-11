package comments

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func CreateComment(articleID, userID uint, replyTo uint, body string) error {
	if _, err := article_utils.CheckPrivateAndExists(userID, articleID); err != nil {
		return err
	}
	return comment_repo.CreateComment(articleID, userID, replyTo, body)
}
