package grpc_article_test

import (
	"testing"

	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
)

func TestGetArticle(t *testing.T) {
	e2e_test.Setup(t)
	articleBeforeEach(t)

}
