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

// Equals checks if another Fund has equal fields to the receiver Fund.
func (f *Fund) Equals(otherFund *Fund) bool {
	IdEq := f.ID == otherFund.ID
	FundCodeEq := f.FundCode == otherFund.FundCode
	SecurityIdEq := f.SecurityID == otherFund.SecurityID
	FundShortNameEq := f.FundShortName == otherFund.FundShortName
	NavDateEq := f.NAVDate == otherFund.NAVDate
	NavEq := f.NAV == otherFund.NAV
	PriorNavEq := f.PriorNAV == otherFund.PriorNAV
	ChangeEq := f.Change == otherFund.Change
	PctChangeEq := f.PctChange == otherFund.PctChange
	SortEq := f.Sort == otherFund.Sort
	YtdReturnEq := f.YtdReturn == otherFund.YtdReturn
	YtdReturnMonthlyEq := f.YtdReturnMonthly == otherFund.YtdReturnMonthly
	YtdReturnQuarterlyEq := f.YtdReturnQuarterly == otherFund.YtdReturnQuarterly
	OneYrReturnEq := f.OneYrReturn == otherFund.OneYrReturn
	OneYrReturnMonthlyEq := f.OneYrReturnMonthly == otherFund.OneYrReturnMonthly
	OneYrReturnQuarterlyEq := f.OneYrReturnQuarterly == otherFund.OneYrReturnQuarterly
	ThreeYrAvgEq := f.ThreeYrAvg == otherFund.ThreeYrAvg
	ThreeYrAvgMonthlyEq := f.ThreeYrAvgMonthly == otherFund.ThreeYrAvgMonthly
	ThreeYrAvgQuarterlyEq := f.ThreeYrAvgQuarterly == otherFund.ThreeYrAvgQuarterly
	FiveYrAvgEq := f.FiveYrAvg == otherFund.FiveYrAvg
	FiveYrAvgMonthlyEq := f.FiveYrAvgMonthly == otherFund.FiveYrAvgMonthly
	FiveYrAvgQuarterlyEq := f.FiveYrAvgQuarterly == otherFund.FiveYrAvgQuarterly
	TenYrAvgEq := f.TenYrAvg == otherFund.TenYrAvg
	TenYrAvgMonthlyEq := f.TenYrAvgMonthly == otherFund.TenYrAvgMonthly
	TenYrAvgQuarterlyEq := f.TenYrAvgQuarterly == otherFund.TenYrAvgQuarterly
	InceptAvgEq := f.InceptAvg == otherFund.InceptAvg
	InceptAvgMonthlyEq := f.InceptAvgMonthly == otherFund.InceptAvgMonthly
	InceptAvgQuarterlyEq := f.InceptAvgQuarterly == otherFund.InceptAvgQuarterly
	SymbolEq := f.Symbol == otherFund.Symbol
	AssetTypeEq := f.AssetType == otherFund.AssetType
	InceptionDateEq := f.InceptionDate == otherFund.InceptionDate
	LegalName2Eq := f.LegalName2 == otherFund.LegalName2
	SeriesNameEq := f.SeriesName == otherFund.SeriesName
	DisplayNameEq := f.DisplayName == otherFund.DisplayName
	DisplayName_Eq := f.DisplayName_ == otherFund.DisplayName_
	CategoryEq := f.Category == otherFund.Category
	AnnualReportEq := f.AnnualReport == otherFund.AnnualReport
	SemiAnnualReportEq := f.SemiAnnualReport == otherFund.SemiAnnualReport
	CusipEq := f.Cusip == otherFund.Cusip
	QuarterlyReportEq := f.QuarterlyReport == otherFund.QuarterlyReport
	ProspectusEq := f.Prospectus == otherFund.Prospectus
	SaiEq := f.Sai == otherFund.Sai
	SoiEq := f.Soi == otherFund.Soi
	FactsheetEq := f.Factsheet == otherFund.Factsheet
	CommentaryEq := f.Commentary == otherFund.Commentary
	LastMonthEndEq := f.LastMonthEnd == otherFund.LastMonthEnd
	LastQtrEnd2Eq := f.LastQtrEnd2 == otherFund.LastQtrEnd2

	return (ChangeEq && PctChangeEq && FundCodeEq && SecurityIdEq && FundShortNameEq && NavDateEq && NavEq && PriorNavEq && SortEq && YtdReturnEq && YtdReturnMonthlyEq && YtdReturnQuarterlyEq && OneYrReturnEq && OneYrReturnMonthlyEq && LastQtrEnd2Eq && LastMonthEndEq && CommentaryEq && FactsheetEq && SoiEq && SaiEq && ProspectusEq && QuarterlyReportEq && CusipEq && SemiAnnualReportEq && AnnualReportEq && CategoryEq && DisplayName_Eq && DisplayNameEq && SeriesNameEq && LegalName2Eq && InceptionDateEq && AssetTypeEq && SymbolEq && InceptAvgQuarterlyEq && InceptAvgMonthlyEq && InceptAvgEq && TenYrAvgQuarterlyEq && TenYrAvgMonthlyEq && TenYrAvgEq && FiveYrAvgQuarterlyEq && FiveYrAvgMonthlyEq && FiveYrAvgEq && ThreeYrAvgQuarterlyEq && ThreeYrAvgMonthlyEq && ThreeYrAvgEq && OneYrReturnQuarterlyEq && IdEq)
}

// A FundsMap represents a map of Fund objects with their symbols as keys.
type FundsMap map[string]Fund
