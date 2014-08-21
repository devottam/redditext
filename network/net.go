package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetches and returns the content of given URL in bytes.
// Explicitly checks for status 200.
func ContentFromURL(url *string) ([]byte, error) {
	res, err := http.Get(*url)
	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch URL: %s", *url)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expecting the response to be 200 but was otherwise: %d", res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}
