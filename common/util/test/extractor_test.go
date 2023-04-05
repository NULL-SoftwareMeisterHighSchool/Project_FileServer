package util_test

import (
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func TestExtractImageURLOK(t *testing.T) {
	// given
	body := []byte(`
	# My Favorite Animal

	![A cute cat](https://www.example.com/cat.jpg)
	
	I absolutely love cats! They are such amazing creatures with their soft fur, cute little paws, and playful personalities. There's nothing quite like curling up on the couch with a purring cat on your lap.
	
	This particular image captures the essence of what I love about cats: their sweet, innocent expressions and their ability to bring joy into our lives. Every time I see a cat like this, I can't help but smile.
	
	If you're a cat lover like me, I'm sure you can relate. And even if you're not, I hope this image brings a little bit of happiness to your day.
	`)

	// when
	imageURLs := util.ExtractImageURL(body)

	// then
	if len(imageURLs) == 0 || imageURLs[0] != "https://www.example.com/cat.jpg" {
		t.Fatalf("extraction failed. given arr: %v", imageURLs)
	}
}

func TestExtractAccessFromHeaderOK(t *testing.T) {
	// given
	header := make(map[string]string)
	token := "hi"
	header["Authorization"] = "Bearer " + token

	// when
	access, err := util.ExtractAccessFromHeader(header)

	// then
	if err != nil || access != token {
		t.Fatalf("invalid token. err: %s, token: %s", err, access)
	}
}

func TestExtractAccessFromHeaderFAIL(t *testing.T) {
	// given
	header := make(map[string]string)
	header["Authorization"] = "Beared wrong-token"

	// when
	_, err := util.ExtractAccessFromHeader(header)

	// then
	if err == nil {
		t.Fatal("error should not be nil")
		return
	}
	if err != errors.ErrAuthFailed {
		t.Fatalf("wrong err type. err: %s", err)
	}
}
