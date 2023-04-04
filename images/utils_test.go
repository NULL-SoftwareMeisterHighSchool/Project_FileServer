package images

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestGetNameAndExtensionOK(t *testing.T) {
	// given
	testImage := "https://domain.com/path/to/image.png"

	// when
	name, extension := getNameAndExtension(testImage)

	// then
	if _, err := uuid.FromString(name); err != nil {
		t.Fatalf("uuid is not valid. given: %s", name)
	}

	if extension != "png" {
		t.Fatalf("extension is not valid. given: %s", extension)
	}
}

func TestCheckExtensionOK(t *testing.T) {
	// given
	allowedExtensions := []string{
		"png", "jpg",
	}
	givenExtension := "png"

	// when & then
	if ok := checkExtension(givenExtension, allowedExtensions); !ok {
		t.Fatal("extension should match")
	}
}

func TestCheckExtensionFAIL(t *testing.T) {
	// given
	allowedExtensions := []string{
		"png", "jpg",
	}
	givenExtension := "gif"

	// when & then
	if ok := checkExtension(givenExtension, allowedExtensions); ok {
		t.Fatal("extension should not match")
	}
}
