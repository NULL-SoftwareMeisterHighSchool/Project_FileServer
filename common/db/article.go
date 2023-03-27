package db

import "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"

type Article struct {
	ID        uint     `json:"id"`
	Summary   string   `json:"summary"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"-"`
}

const MAX_SUMMARY_LENGTH = 100

func (a *Article) SetSummaryFromBody(body []byte) {
	plainText := util.SanitizeExceptPlainText(body)
	if len(plainText) < MAX_SUMMARY_LENGTH {
		a.Summary = string(plainText)
	} else {
		a.Summary = string(plainText[:MAX_SUMMARY_LENGTH])
	}
}

func (a *Article) SetImagesFromBody(body []byte) {

}

func (a *Article) SetThumbnail() {

}
