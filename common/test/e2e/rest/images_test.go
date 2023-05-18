package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	article_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/repo"
	user_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/repo"
	image_controller "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest/server/images"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
	"github.com/gofiber/fiber/v2"
)

func TestUploadImage(t *testing.T) {

	const userID uint = 1

	// Setup
	e2e_test.Setup(t)

	if err := user_repo.CreateUserByID(userID); err != nil {
		t.Error(err)
	}
	articleID, err := article_repo.CreateArticle(userID, "", article_entity.TYPE_GENERAL)
	if err != nil {
		t.Error(err)
	}

	app := fiber.New()
	app.Post("/images",
		// fake userID
		func(c *fiber.Ctx) error {
			c.Locals("userID", userID)
			return c.Next()
		},
		image_controller.UploadImage,
	)

	// Create MultipartForm
	file, _ := os.Open("../resources/unknown.png")
	defer file.Close()
	buf := &bytes.Buffer{}

	w := multipart.NewWriter(buf)
	part, _ := w.CreateFormFile("image", "unknown.png")
	img, _ := png.Decode(file)
	png.Encode(part, img)
	w.Close()

	// Create Request
	req := httptest.NewRequest("POST", "/images?articleID="+fmt.Sprint(articleID), buf)
	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
		return
	}
	// Read Response
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		if err != nil {
			t.Error(err)
		}
		t.Error("wtf it didn't accept:", string(b))
		return
	}

	var url struct{ URL string }
	if err := json.Unmarshal(b, &url); err != nil {
		t.Error(err)
		return
	}

	// s3
	// res, err = http.DefaultClient.Get(url.URL)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// if res.StatusCode != http.StatusOK {
	// 	t.Errorf("no image found on path: %s. status: %d", url.URL, res.StatusCode)
	// 	return
	// }

	// filesystem
	if _, err := os.Stat("." + strings.Split(url.URL, ":8080")[1]); err != nil {
		t.Error(err)
	}
}
