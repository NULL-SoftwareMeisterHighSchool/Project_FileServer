package util

import (
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var xss = bluemonday.UGCPolicy()
var strict = bluemonday.StrictPolicy()

func SanitizeXSS(body []byte) []byte {
	return xss.SanitizeBytes(body)
}

func SanitizeExceptPlainText(body []byte) []byte {
	withoutHTML := strict.SanitizeBytes(body)
	r := regexp.MustCompile(strings.Join([]string{
		`(#{1,6}\s)(.*)`,                           // heading
		`(\*|\_)+(\S+?)(\*|\_)+`,                   // bold and italic text....
		`(!)?(\[.*\])(\((http)(?:s)?(\:\/\/).*\))`, // links
		"(\\`{1})(.*)(\\`{1})",                     // inline code
		"(\\`{3}\\n+)(.*)(\\n+\\`{3})",             // code block
		`(\=|\-|\*){3}`,                            // horizontal line
		`(((\|)([a-zA-Z\d+\s#!@'"():;\\\/.\[\]\^<={$}>?(?!-))]+))+(\|))(?:\n)?((\|)(-+))+(\|)(\n)((\|)(\W+|\w+|\S+))+(\|$)`, // TABLE
	}, "|"))
	return r.ReplaceAll(withoutHTML, []byte(""))
}
