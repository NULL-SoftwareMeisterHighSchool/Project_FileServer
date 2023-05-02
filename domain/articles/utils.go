package articles

import (
	"strings"

	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
)

func filterDeletableImageURLByEndpoint(urls []string, endpoint string) []string {
	shouldDelete := []string{}
	for _, url := range urls {
		if strings.HasPrefix(url, endpoint) {
			shouldDelete = append(shouldDelete, url)
		}
	}
	return shouldDelete
}

func checkPrivateAndExists(userID, articleID uint) error {
	pInfo, err := getInfoOrNotFound(articleID)
	if err != nil {
		return err
	}

	if pInfo.IsPrivate && pInfo.AuthorID != userID {
		return article_errors.ErrPermissionDenied
	}
	return nil
}

func checkSudoAndExists(userID, articleID uint) error {
	pInfo, err := getInfoOrNotFound(articleID)
	if err != nil {
		return err
	}

	if pInfo.AuthorID != userID {
		return article_errors.ErrPermissionDenied
	}
	return nil
}

func checkExists(articleID uint) error {
	_, err := getInfoOrNotFound(articleID)
	return err
}

func getInfoOrNotFound(articleID uint) (*article_repo.ArticlePermissionInfo, error) {
	pInfo, err := article_repo.GetArticlePermissionInfoByID(articleID)
	if err != nil {
		return nil, article_errors.ErrNotFound
	}

	return pInfo, nil
}
