package e2e_test

import (
	"os"
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
	image_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/entity"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
)

func tearDown(t *testing.T) {
	if err := database.DB().Migrator().DropTable(
		&article_entity.Article{},
		&article_entity.ArticleBody{},
		&comment_entity.Comment{},
		&image_entity.Image{},
		&user_entity.User{},
		"user_likes",
	); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}

	// if storage uses filesystem
	os.RemoveAll("./contents")
}
