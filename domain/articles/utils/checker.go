package article_utils

import (
	"errors"

	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
	"gorm.io/gorm"
)

func CheckPrivateAndExists(userID, articleID uint) (isAuthor bool, err error) {
	pInfo, err := getInfoOrNotFound(articleID)
	if err != nil {
		return false, err
	}

	isAuthor = pInfo.AuthorID == userID

	if pInfo.IsPrivate && !isAuthor {
		return isAuthor, article_errors.ErrPermissionDenied
	}
	return isAuthor, nil
}

func CheckSudoAndExists(userID, articleID uint) error {
	pInfo, err := getInfoOrNotFound(articleID)
	if err != nil {
		return err
	}

	if pInfo.AuthorID != userID {
		return article_errors.ErrPermissionDenied
	}
	return nil
}

func CheckExists(articleID uint) error {
	_, err := getInfoOrNotFound(articleID)
	return err
}

func getInfoOrNotFound(articleID uint) (*article_repo.ArticlePermissionInfo, error) {
	pInfo, err := article_repo.GetArticlePermissionInfoByID(articleID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, article_errors.ErrArticleNotFound
		}
		return nil, err
	}

	return pInfo, nil
}
