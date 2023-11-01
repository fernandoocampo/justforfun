package numbers_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/numbers"
)

func TestFakeBin(t *testing.T) {
	// Given
	cases := map[string]struct {
		input string
		want  string
	}{
		"45385593107843568": {
			input: "45385593107843568",
			want:  "01011110001100111",
		},
		"509321967506747": {
			input: "509321967506747",
			want:  "101000111101101",
		},
		"366058562030849490134388085": {
			input: "366058562030849490134388085",
			want:  "011011110000101010000011011",
		},
		"15889923": {
			input: "15889923",
			want:  "01111100",
		},
		"800857237867": {
			input: "800857237867",
			want:  "100111001111",
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// When
			got := numbers.FakeBin(testData.input)
			// Then
			if testData.want != got {
				t.Errorf("want: %s, but got: %s", testData.want, got)
			}
		})
	}
}
