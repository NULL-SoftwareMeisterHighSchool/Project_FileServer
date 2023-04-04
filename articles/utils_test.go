package articles

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

func TestGetIdsFromQueryOK(t *testing.T) {
	// given
	idArrStr := "12,31,50,112"

	// when
	ids := getIdsFromQuery(idArrStr)

	// then
	if len(ids) == 0 {
		t.Fatal("ids should be list of uints")
	}

	if ids[0] != 12 {
		t.Fatalf("parsing failed. given arr: %v", ids)
	}
}

func TestGetIdsFromQueryFAIL(t *testing.T) {
	// given
	idArrStr := "-12,31,50 ,hi"

	// when
	ids := getIdsFromQuery(idArrStr)

	// then
	if len(ids) != 0 {
		t.Fatalf("ids should be empty. given arr: %v", ids)
	}
}
