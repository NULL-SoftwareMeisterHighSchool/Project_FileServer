package article_utils

import "testing"

func TestFilterDeletableImageURLByEndpointOK(t *testing.T) {
	// given
	correctEndpoint := "https://correct.endpoint/images"
	urls := []string{
		correctEndpoint + "/indeed",
		"https://incorrect.endpoint",
	}

	// when
	filtered := filterDeletableImageURLByEndpoint(urls, correctEndpoint)

	// then
	if len(filtered) == 0 || filtered[0] != correctEndpoint+"/indeed" {
		t.Fatalf("filtering failed. given arr: %v", filtered)
	}
}

func TestFilterDeletableImageURLByEndpointFAIL(t *testing.T) {
	// given
	correctEndpoint := "https://correct.endpoint/images"
	urls := []string{
		"https://nicorrect.endpointy/indeed",
		"https://incorrect.endpoint",
	}

	// when
	filtered := filterDeletableImageURLByEndpoint(urls, correctEndpoint)

	// then
	if len(filtered) != 0 {
		t.Fatalf("filtered array should be empty. given: %v", filtered)
	}
}
