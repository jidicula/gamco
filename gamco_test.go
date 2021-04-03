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
