package grpc_article_test

import (
	"context"
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

	techArticleID, _ := article_repo.CreateArticle(anotherUserID, "", article_entity.TYPE_TECH)
	article_repo.UpdateIsPrivateByID(techArticleID, false)

	titledArticleID, _ := article_repo.CreateArticle(userID, "what's up kimwonwuk", article_entity.TYPE_GENERAL)
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

	// only general
	list, _ := client.ListArticle(context.Background(), &articles_pb.ListArticleRequest{
		Offset: 0,
		Amount: 10,
		Order:  articles_pb.ListArticleOrder_POPULARITY,
		Type:   articles_pb.ArticleType_GENERAL,
	})
	if len(list.GetArticles()) != 2 {
		t.Errorf("expected: %d articles. got: %+#v", 2, list.GetArticles())
	} else if articleType := list.GetArticles()[0].Type; articleType != articles_pb.ArticleType_GENERAL {
		t.Errorf("wrong article type. expected: %d, got: %d", articles_pb.ArticleType_GENERAL, articleType)
	}

	// by user
	aUserID := uint32(anotherUserID)
	list, _ = client.ListArticle(context.Background(), &articles_pb.ListArticleRequest{
		Offset:   0,
		Amount:   10,
		AuthorID: &aUserID,
	})
	if len(list.GetArticles()) != 2 {
		t.Errorf("expected: %d articles. got: %+#v", 2, list.GetArticles())
	} else if authorID := list.GetArticles()[0].AuthorID; authorID != aUserID {
		t.Errorf("wrong article type. expected: %d, got: %d", aUserID, authorID)
	}

	// query for title
	query := "kimwonwuk"
	list, _ = client.ListArticle(context.Background(), &articles_pb.ListArticleRequest{
		Offset: 0,
		Amount: 10,
		Query:  &query,
	})
	if len(list.GetArticles()) != 1 {
		t.Errorf("expected: %d articles. got: %+#v", 1, list.GetArticles())
	} else if id := list.GetArticles()[0].ArticleID; id != uint32(titledArticleID) {
		t.Errorf("wrong article type. expected: %d, got: %d", titledArticleID, id)
	}

	// query for content
	query = "ipsum"
	list, _ = client.ListArticle(context.Background(), &articles_pb.ListArticleRequest{
		Offset: 0,
		Amount: 10,
		Query:  &query,
	})
	if len(list.GetArticles()) != 1 {
		t.Errorf("expected: %d article. got: %+#v", 1, list.GetArticles())
	} else if id := list.GetArticles()[0].ArticleID; id != uint32(fulltextArticleID) {
		t.Errorf("wrong article type. expected: %d, got: %d", fulltextArticleID, id)
	}
}
