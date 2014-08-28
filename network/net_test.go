package network

import (
	"testing"
)

func TestContentFromExistingURL(t *testing.T) {
	url := "https://godoc.org/testing"
	b, err := ContentFromURL(&url)

	if len(b) < 10 || err != nil {
		t.Errorf("Failed to fetch the given URL: %v. \nReturned bytes: %v", url, string(b))
	}
}
