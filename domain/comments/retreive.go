package comments

import (
	comment_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/repo"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func GetCommentsFromArticle(userID, articleID uint) ([]*comment_repo.CommentWithReplyCount, error) {

	if _, err := article_utils.CheckPrivateAndExists(userID, articleID); err != nil {
		return nil, err
	}

	comments, err := comment_repo.GetCommentsByArticleID(articleID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
