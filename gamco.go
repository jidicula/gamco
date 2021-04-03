// Copyright (C) 2021  Johanan Idicula
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
