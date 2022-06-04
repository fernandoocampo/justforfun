package vowels_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/vowels"
)

func TestGetCountFor(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"abracadabra": {
			input: "abracadabra",
			want:  5,
		},
	}
	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := vowels.GetCountFor(testdata.input)
			if got != testdata.want {
				subt.Errorf("want: %d, but got: %d", testdata.want, got)
			}
		})
	}
}

func TestGetCountRec(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"abracadabra": {
			input: "abracadabra",
			want:  5,
		},
	}
	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := vowels.GetCountRec(testdata.input)
			if got != testdata.want {
				subt.Errorf("want: %d, but got: %d", testdata.want, got)
			}
		})
	}
}
