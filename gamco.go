// Package gamco provides an unofficial API wrapper for GAMCO's Closed-End
// Funds API
package gamco

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// getData hits the nav_closed_ends endpoint.
func getData() (string, error) {
	url := "https://gabdotcom-api.com/api/v1/nav_closed_ends"
	var bodyBytes []byte
	var bodyString string
	c := http.Client{}

	resp, err := c.Get(url)
	if err != nil {
		return bodyString, fmt.Errorf("HTTP GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return bodyString, fmt.Errorf("API call failed, response status %v", resp.StatusCode)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return bodyString, fmt.Errorf("Response decoding failed: %v", err)
	}
	bodyString = string(bodyBytes)

	return bodyString, err
}
