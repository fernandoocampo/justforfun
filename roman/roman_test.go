package roman_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/roman"
)

func TestGetRomanNumbers(t *testing.T) {
	cases := map[string]struct {
		givenNumber int
		want        string
	}{
		"1_I": {
			givenNumber: 1,
			want:        "I",
		},
		"2_II": {
			givenNumber: 2,
			want:        "II",
		},
		"3_III": {
			givenNumber: 3,
			want:        "III",
		},
		"4_IV": {
			givenNumber: 4,
			want:        "IV",
		},
		"5_V": {
			givenNumber: 5,
			want:        "V",
		},
		"6_VI": {
			givenNumber: 6,
			want:        "VI",
		},
		"73_LXXIII": {
			givenNumber: 73,
			want:        "LXXIII",
		},
		"45_XLV": {
			givenNumber: 45,
			want:        "XLV",
		},
		"445_XLDV": {
			givenNumber: 445,
			want:        "CDXLV",
		},
		"93_XCIII": {
			givenNumber: 93,
			want:        "XCIII",
		},
	}
	for name, testdata := range cases {
		t.Run(name, func(subtest *testing.T) {
			got := roman.GetRomanNumbers(testdata.givenNumber)
			if testdata.want != got {
				subtest.Errorf("want: %s, but got %s", testdata.want, got)
			}
		})
	}
}

func TestGetNumberFromRoman(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"I_1": {
			input: "I",
			want:  1,
		},
		"II_2": {
			input: "II",
			want:  2,
		},
		"III_3": {
			input: "III",
			want:  3,
		},
		"IV_4": {
			input: "IV",
			want:  4,
		},
		"VIII_8": {
			input: "VIII",
			want:  8,
		},
		"XX_20": {
			input: "XX",
			want:  20,
		},
		"CDXLV_445": {
			input: "CDXLV",
			want:  445,
		},
		"XXIII_23": {
			input: "XXIII",
			want:  23,
		},
		"XCIII_93": {
			input: "XCIII",
			want:  93,
		},
	}
	for name, testdata := range cases {
		t.Run(name, func(subt *testing.T) {
			got := roman.GetNumberFromRoman(testdata.input)
			if testdata.want != got {
				subt.Errorf("want: %d, but got: %d", testdata.want, got)
			}
		})
	}
}
