package util

import "regexp"

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
