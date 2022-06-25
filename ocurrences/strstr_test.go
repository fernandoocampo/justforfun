package ocurrences_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/ocurrences"
)

func TestStrStr(t *testing.T) {
	cases := map[string]struct {
		haystack string
		needle   string
		want     int
	}{
		"a_a": {
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		"hello_ll": {
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		"aaaaa_bba": {
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		"bbbbababbbaabbba_abb": {
			haystack: "bbbbababbbaabbba",
			needle:   "abb",
			want:     6,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := ocurrences.StrStr(data.haystack, data.needle)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}
