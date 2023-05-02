package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func UpdateArticleBody(articleID, userID uint, body []byte) error {
	if err := checkSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateArticleBodyAndImages(articleID, body, util.ExtractImageURL(body))
	return nil
}

func UpdateArticleTitle(articleID, userID uint, title string) error {
	if err := checkSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateTitleByID(articleID, title)
	return nil
}

func UpdateArticleVisibility(userID, articleID uint, isPrivate bool) error {
	if err := checkSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateIsPrivateByID(articleID, isPrivate)
	return nil
}

func ToggleArticleLike(userID, articleID uint) error {
	if err := checkExists(articleID); err != nil {
		return err
	}

	go article_repo.ToggleLike(articleID, userID)

	return nil
}
