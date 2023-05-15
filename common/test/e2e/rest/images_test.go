package rest_test

import (
	"encoding/json"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	image_controller "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/rest/server/images"
	e2e_test "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/test/e2e"
	"github.com/gofiber/fiber/v2"
)

func TestUploadImage(t *testing.T) {

	// Setup
	e2e_test.Setup(t)

	app := fiber.New()
	app.Post("/images",
		// fake userID
		func(c *fiber.Ctx) error {
			c.Locals("userID", 1)
			return nil
		},
		image_controller.UploadImage,
	)

	// MultipartForm Setup
	pr, pw := io.Pipe()
	w := multipart.NewWriter(pw)

	go func() {
		defer w.Close()
		part, _ := w.CreateFormFile("image", "unkown.png")
		file, _ := os.Open("resources/unknown.png")
		img, _ := png.Decode(file)
		png.Encode(part, img)
	}()

	// Create Request
	req := httptest.NewRequest("POST", "/images?articleID=1", pr)
	req.Header.Add("Content-Type", w.FormDataContentType())

	res, err := app.Test(req, 1)
	if err != nil {
		t.Error(err)
	}

	// Read Response
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		if err != nil {
			t.Error(err)
		}
		t.Error("wtf it didn't accept", string(b))
	}

	var url struct{ URL string }
	if err := json.Unmarshal(b, &url); err != nil {
		t.Error(err)
	}

	res, err = http.DefaultClient.Get(url.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("no image found on path: %s. status: %d", url.URL, res.StatusCode)
	}

	// Tear Down
	e2e_test.TearDown(t)
}
