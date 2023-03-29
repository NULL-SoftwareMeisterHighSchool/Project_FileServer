package util

import (
	"regexp"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/errors"
)

var imageLinkRegex = regexp.MustCompile(`(\!)?(\[.*\])(\((http)(?:s)?(\:\/\/).*(\.(jpg|png|gif|tiff|bmp))(?:(\s\"|\')(\w|\W|\d)*(\"|\'))?\))`)
var imageRegex = regexp.MustCompile(`(http)(?:s)?(\:\/\/).*(\.(jpg|png|gif|tiff|bmp))`)

func ExtractImageURL(body []byte) []string {
	extracted := []string{}
	matchingLinks := imageLinkRegex.FindAll(body, -1)

	for _, link := range matchingLinks {
		extracted = append(extracted, string(imageRegex.Find(link)))
	}

	return extracted
}

func ExtractAccessFromHeader(header map[string]string) (string, error) {
	tokenRaw := header["Authorization"]
	tokenArr := strings.Split(tokenRaw, " ")
	if tokenArr[0] != "Bearer" || tokenArr[1] == "" {
		return "", errors.AuthFailedError
	}
	return tokenArr[1], nil
}
