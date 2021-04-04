// github.com/jidicula/go-gamco provides an unofficial API wrapper for GAMCO's
// Closed-End Funds API.
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
	"time"
)

// getData hits the nav_closed_ends endpoint and stringifies the response.
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

// A Fund represents a single closed-end GAMCO fund.
type Fund struct {
	ID                   int       `json:"id"`
	FundCode             int       `json:"fund_code"`
	SecurityID           string    `json:"security_id"`
	FundShortName        string    `json:"fundshortname"`
	NAVDate              time.Time `json:"pricedate"`
	NAV                  string    `json:"price"`
	PriorNAV             string    `json:"prior_price"`
	Change               string    `json:"change"`
	PctChange            string    `json:"pct_change"`
	Sort                 string    `json:"sort"`
	YtdReturn            float64   `json:"ytd_return"`
	YtdReturnMonthly     float64   `json:"ytd_return_monthly"`
	YtdReturnQuarterly   float64   `json:"ytd_return_quarterly"`
	OneYrReturn          float64   `json:"one_yr_return"`
	OneYrReturnMonthly   float64   `json:"one_yr_return_monthly"`
	OneYrReturnQuarterly float64   `json:"one_yr_return_quarterly"`
	ThreeYrAvg           float64   `json:"three_yr_avg"`
	ThreeYrAvgMonthly    float64   `json:"three_yr_avg_monthly"`
	ThreeYrAvgQuarterly  float64   `json:"three_yr_avg_quarterly"`
	FiveYrAvg            float64   `json:"five_yr_avg"`
	FiveYrAvgMonthly     float64   `json:"five_yr_avg_monthly"`
	FiveYrAvgQuarterly   float64   `json:"five_yr_avg_quarterly"`
	TenYrAvg             float64   `json:"ten_yr_avg"`
	TenYrAvgMonthly      float64   `json:"ten_yr_avg_monthly"`
	TenYrAvgQuarterly    float64   `json:"ten_yr_avg_quarterly"`
	InceptAvg            float64   `json:"incept_avg"`
	InceptAvgMonthly     float64   `json:"incept_avg_monthly"`
	InceptAvgQuarterly   float64   `json:"incept_avg_quarterly"`
	Symbol               string    `json:"symbol"`
	AssetType            string    `json:"asset_type"`
	InceptionDate        time.Time `json:"inception_date"`
	LegalName2           string    `json:"legalname2"`
	SeriesName           string    `json:"seriesname"`
	DisplayName          string    `json:"displayname"`
	DisplayName_         string    `json:"displayname_"`
	Category             string    `json:"category"`
	AnnualReport         string    `json:"annual_report"`
	SemiAnnualReport     string    `json:"semi_annual_report"`
	Cusip                string    `json:"cusip"`
	QuarterlyReport      string    `json:"quarterly_report"`
	Prospectus           string    `json:"prospectus"`
	Sai                  string    `json:"sai"`
	Soi                  string    `json:"soi"`
	Factsheet            string    `json:"factsheet"`
	Commentary           string    `json:"commentary"`
	LastMonthEnd         time.Time `json:"last_month_end"`
	LastQtrEnd2          time.Time `json:"last_qtr_end_2"`
}
