package articles

import (
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
	article_utils "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/utils"
)

func UpdateArticleBody(articleID, userID uint, body []byte) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateArticleBodyAndImages(articleID, body, util.ExtractImageURL(body))
	return nil
}

func UpdateArticleTitle(articleID, userID uint, title string) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateTitleByID(articleID, title)
	return nil
}

func UpdateArticleVisibility(userID, articleID uint, isPrivate bool) error {
	if err := article_utils.CheckSudoAndExists(userID, articleID); err != nil {
		return err
	}

	go article_repo.UpdateIsPrivateByID(articleID, isPrivate)
	return nil
}

func ToggleArticleLike(userID, articleID uint) error {
	if err := article_utils.CheckExists(articleID); err != nil {
		return err
	}

	go article_repo.ToggleLike(articleID, userID)

	return nil
}
