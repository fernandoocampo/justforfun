package pianos_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/pianos"
)

func TestGuessColorKey(t *testing.T) {
	cases := map[string]struct {
		input int
		want  int
	}{
		"3": {
			input: 3,
			want:  3,
		},
		"56": {
			input: 56,
			want:  3,
		},
		"63": {
			input: 63,
			want:  10,
		},
	}

	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := pianos.GuestColorKey(testdata.input)
			if got != testdata.want {
				subt.Errorf("want: %d, but got %d", testdata.want, got)
			}
		})
	}
}
