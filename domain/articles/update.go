package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
)

func UpdateArticleBody(articleID, userID uint, body []byte) error {

	pInfo, err := article_repo.GetArticlePermissionInfoByID(articleID)
	if err != nil {
		return article_errors.ErrNotFound
	}
	if pInfo.AuthorID != userID {
		return article_errors.ErrPermissionDenied
	}

	// do your thing...

	return nil
}
