package db

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"

const MAX_SUMMARY_LENGTH = 100

type Article struct {
	ID        uint     `json:"id"`
	Summary   string   `json:"summary"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"-"`
}

func (a *Article) SetSummaryFromBody(body []byte) {
	plainText := util.SanitizeExceptPlainText(body)
	minLength := util.GetMin(len(plainText), MAX_SUMMARY_LENGTH)
	a.Summary = string(plainText[:minLength])
}

func (a *Article) SetImagesFromBody(body []byte) {
	a.Images = util.ExtractImageURL(body)
}

func (a *Article) SetThumbnail() {
	if len(a.Images) == 0 {
		a.Thumbnail = ""
	} else {
		a.Thumbnail = a.Images[0]
	}
}
