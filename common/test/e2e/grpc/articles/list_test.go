package grpc_article_test

import (
	"context"
	"log"
	"testing"

	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	user_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/repo"
	articles_pb "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/grpc/server/pb/articles"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
)

func TestListArticles(t *testing.T) {
	// given
	e2e_test.Setup(t)
	client := articleBeforeEach(t)

	const (
		userID        uint = 1
		anotherUserID uint = 2
		articleCount       = 4
	)

	user_repo.CreateUserByID(userID)
	user_repo.CreateUserByID(anotherUserID)

	generalArticleID, _ := article_repo.CreateArticle(userID, "", article_entity.TYPE_GENERAL)
	article_repo.UpdateIsPrivateByID(generalArticleID, false)

	techArticleID, _ := article_repo.CreateArticle(userID, "", article_entity.TYPE_TECH)
	article_repo.UpdateIsPrivateByID(techArticleID, false)

	titledArticleID, _ := article_repo.CreateArticle(userID, "what's up", article_entity.TYPE_GENERAL)
	article_repo.UpdateIsPrivateByID(titledArticleID, false)

	fulltextArticleID, _ := article_repo.CreateArticle(anotherUserID, "this is fulltext", article_entity.TYPE_TECH)
	article_repo.UpdateIsPrivateByID(fulltextArticleID, false)
	article_repo.UpdateArticleBody(fulltextArticleID, []byte(`
	Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
	Quisque et dapibus lorem, in elementum massa. 
	Nam eget purus volutpat, vulputate nisi aliquet, vehicula eros. 
	Donec porta eu nulla at dictum. 
	Suspendisse eleifend eros vitae nisl bibendum, eget fringilla velit bibendum. 
	Quisque finibus neque quis risus eleifend egestas. 
	Aenean felis odio, lacinia in porta quis, commodo a turpis. 
	Vivamus blandit tempus aliquet. Cras ac mollis arcu, sed luctus magna. 
	Maecenas blandit ligula et odio tristique, id porta lectus volutpat
	`))

	// default
	if list, _ := client.ListArticle(context.Background(), &articles_pb.ListArticleRequest{
		Offset: 0,
		Amount: 10,
		Order:  articles_pb.ListArticleOrder_POPULARITY,
	}); len(list.GetArticles()) != articleCount {
		t.Errorf("expected: %d articles. got: %+#v", articleCount, list.GetArticles())
	}

	log.Println(generalArticleID, techArticleID, titledArticleID)
}
