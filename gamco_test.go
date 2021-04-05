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

package gamco

import (
	"encoding/json"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestGetData(t *testing.T) {
	tests := map[string]struct {
		want *regexp.Regexp
	}{
		"api test": {
			want: regexp.MustCompile(`\[(\{(\".*\":.*,?)+\},?)+\]`),
		},
	}

	for name, tt := range tests {

		t.Run(name, func(t *testing.T) {

			got, err := getData()
			if err != nil {
				t.Fatalf(err.Error())
			}
			if !tt.want.MatchString(got) {
				t.Errorf("%s: got %s, want `%s` regex match", name, got, tt.want.String())
			}
		})
	}
}

func TestFundUnmarshal(t *testing.T) {
	// Date setup
	dates, err := dateSetup("2021-04-01T00:00:00.000Z", "1999-07-09T00:00:00.000Z", "03/31/2021", "03/31/2021")
	if err != nil {
		t.Fatalf(err.Error())
	}
	priceDate := dates["priceDate"]
	inceptionDate := dates["inceptionDate"]
	lastMonthEnd := dates["lastMonthEnd"]
	lastQtrEnd := dates["lastQtrEnd"]

	tests := map[string]struct {
		data []byte
		want Fund
	}{
		"fund": {data: []byte(`{
    "id": 515,
    "fund_code": -113,
    "security_id": "36240A101",
    "fundshortname": "Utility Trust",
    "pricedate": "2021-04-01T00:00:00.000Z",
    "price": "4.27",
    "prior_price": "4.25",
    "change": "0.02",
    "pct_change": "0.004706",
    "sort": "43.0",
    "ytd_return": 0.0767238547,
    "ytd_return_monthly": 0.0716806516,
    "ytd_return_quarterly": 0.0716806516,
    "one_yr_return": 0.3799871231,
    "one_yr_return_monthly": 0.2908244253,
    "one_yr_return_quarterly": 0.2908244253,
    "three_yr_avg": 0.0821346464,
    "three_yr_avg_monthly": 0.0806954728,
    "three_yr_avg_quarterly": 0.0806954728,
    "five_yr_avg": 0.0628642424,
    "five_yr_avg_monthly": 0.0618578521,
    "five_yr_avg_quarterly": 0.0618578521,
    "ten_yr_avg": 0.0852284004,
    "ten_yr_avg_monthly": 0.086077124,
    "ten_yr_avg_quarterly": 0.086077124,
    "incept_avg": 0.0851655141,
    "incept_avg_monthly": 0.085533555,
    "incept_avg_quarterly": 0.0869319686,
    "symbol": "GUT",
    "asset_type": "Equity",
    "inception_date": "1999-07-09T00:00:00.000Z",
    "legalname2": "The Gabelli Utility Trust",
    "seriesname": null,
    "displayname": "Gabelli Utility Trust",
    "displayname_": "The Gabelli Utility Trust",
    "category": "value",
    "annual_report": "https://gab-annual-reports.s3.us-east-2.amazonaws.com/GUTFundWebReady12312020.pdf",
    "semi_annual_report": "https://gab-semi-annuals.s3.us-east-2.amazonaws.com/TheGabelliUtilityTrust606302020.pdf",
    "cusip": "36240A101",
    "quarterly_report": "https://gab-reports.s3.us-east-2.amazonaws.com/2006q3/-113.pdf",
    "prospectus": "https://gab-prospectus.s3.us-east-2.amazonaws.com/-113.pdf",
    "sai": "https://gab-sai.s3.us-east-2.amazonaws.com/-113_sai.pdf",
    "soi": null,
    "factsheet": "https://gab-factsheets.s3.us-east-2.amazonaws.com/closedEnd_FactSheets4Q2020DRAFT_GUT12312020.pdf",
    "commentary": "https://gab-commentary-pdf.s3.us-east-2.amazonaws.com/WEB_CEF_4Q2012312020.pdf",
    "last_month_end": "03/31/2021",
    "last_qtr_end_2": "03/31/2021"
  }`), want: Fund{
			ID:                   515,
			FundCode:             -113,
			SecurityID:           "36240A101",
			FundShortName:        "Utility Trust",
			NAVDate:              priceDate,
			NAV:                  "4.27",
			PriorNAV:             "4.25",
			Change:               "0.02",
			PctChange:            "0.004706",
			Sort:                 "43.0",
			YtdReturn:            0.0767238547,
			YtdReturnMonthly:     0.0716806516,
			YtdReturnQuarterly:   0.0716806516,
			OneYrReturn:          0.3799871231,
			OneYrReturnMonthly:   0.2908244253,
			OneYrReturnQuarterly: 0.2908244253,
			ThreeYrAvg:           0.0821346464,
			ThreeYrAvgMonthly:    0.0806954728,
			ThreeYrAvgQuarterly:  0.0806954728,
			FiveYrAvg:            0.0628642424,
			FiveYrAvgMonthly:     0.0618578521,
			FiveYrAvgQuarterly:   0.0618578521,
			TenYrAvg:             0.0852284004,
			TenYrAvgMonthly:      0.086077124,
			TenYrAvgQuarterly:    0.086077124,
			InceptAvg:            0.0851655141,
			InceptAvgMonthly:     0.085533555,
			InceptAvgQuarterly:   0.0869319686,
			Symbol:               "GUT",
			AssetType:            "Equity",
			InceptionDate:        inceptionDate,
			LegalName2:           "The Gabelli Utility Trust",
			SeriesName:           "",
			DisplayName:          "Gabelli Utility Trust",
			DisplayName_:         "The Gabelli Utility Trust",
			Category:             "value",
			AnnualReport:         "https://gab-annual-reports.s3.us-east-2.amazonaws.com/GUTFundWebReady12312020.pdf",
			SemiAnnualReport:     "https://gab-semi-annuals.s3.us-east-2.amazonaws.com/TheGabelliUtilityTrust606302020.pdf",
			Cusip:                "36240A101",
			QuarterlyReport:      "https://gab-reports.s3.us-east-2.amazonaws.com/2006q3/-113.pdf",
			Prospectus:           "https://gab-prospectus.s3.us-east-2.amazonaws.com/-113.pdf",
			Sai:                  "https://gab-sai.s3.us-east-2.amazonaws.com/-113_sai.pdf",
			Soi:                  "",
			Factsheet:            "https://gab-factsheets.s3.us-east-2.amazonaws.com/closedEnd_FactSheets4Q2020DRAFT_GUT12312020.pdf",
			Commentary:           "https://gab-commentary-pdf.s3.us-east-2.amazonaws.com/WEB_CEF_4Q2012312020.pdf",
			LastMonthEnd:         lastMonthEnd,
			LastQtrEnd2:          lastQtrEnd,
		}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Fund{}
			err := json.Unmarshal(tt.data, &got)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}

}

// dateSetup sets up a map of times for use in tests
func dateSetup(priceDate string, inceptionDate string, lastMonthEnd string, lastQtrEnd string) (map[string]time.Time, error) {
	dates := make(map[string]time.Time)
	dateFormat := "01/02/2006"
	var err error
	dates["priceDate"], err = time.Parse(time.RFC3339, priceDate)
	if err != nil {
		return dates, err
	}
	dates["inceptionDate"], err = time.Parse(time.RFC3339, inceptionDate)
	if err != nil {
		return dates, err
	}
	dates["lastMonthEnd"], err = time.Parse(dateFormat, lastMonthEnd)
	if err != nil {
		return dates, err
	}
	dates["lastQtrEnd"], err = time.Parse(dateFormat, lastQtrEnd)
	if err != nil {
		return dates, err
	}

	return dates, err
}
