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
	"regexp"
	"testing"
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
