package util_test

import (
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func TestSanitizeXSS(t *testing.T) {
	// given
	body := []byte(`
	<a href="javascript:alert('XSS')">XSS</a>
	`)

	// when
	sanitized := util.SanitizeXSS(body)

	// then
	t.Log(string(sanitized))
}

func TestSanitizeExceptPlainText(t *testing.T) {
	// given
	body := []byte(`
	# Welcome to my Blog

	<p>This is a blog post about nothing in particular.</p>
	
	## Heading 2
	
	<p>Here's a list of random things:</p>
	
	<ul>
	  <li>Apples</li>
	  <li>Bananas</li>
	  <li>Chickens</li>
	</ul>
	
	<p>And here's a table:</p>
	
	<table>
	  <thead>
		<tr>
		  <th>Column 1</th>
		  <th>Column 2</th>
		</tr>
	  </thead>
	  <tbody>
		<tr>
		  <td>Row 1, Column 1</td>
		  <td>Row 1, Column 2</td>
		</tr>
		<tr>
		  <td>Row 2, Column 1</td>
		  <td>Row 2, Column 2</td>
		</tr>
	  </tbody>
	</table>
	
	<p>Thanks for reading!</p>
	`)

	// when
	sanitized := util.SanitizeExceptPlainText(body)

	// then
	t.Log(string(sanitized))
}
