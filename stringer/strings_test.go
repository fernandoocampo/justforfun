package stringer_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/justforfun/stringer"
)

func TestIsUpperCase(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]bool{
		"c":                      false,
		"C":                      true,
		"hello I AM DONALD":      false,
		"HELLO I AM DONALD":      true,
		"ACSKLDFJSgSKLDFJSKLDFJ": false,
		"ACSKLDFJSGSKLDFJSKLDFJ": true,
	}

	for input, want := range cases {
		t.Run(input, func(st *testing.T) {
			givenValue := input
			wantedValue := want
			st.Parallel()
			// When
			got := stringer.IsUpperCase(givenValue)
			// Then
			if got != wantedValue {
				st.Errorf("%q: want: %t, but got: %t", givenValue, wantedValue, got)
			}
		})
	}
}

func TestIsUpperCaseSmart(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]bool{
		"c":                      false,
		"C":                      true,
		"hello I AM DONALD":      false,
		"HELLO I AM DONALD":      true,
		"ACSKLDFJSgSKLDFJSKLDFJ": false,
		"ACSKLDFJSGSKLDFJSKLDFJ": true,
	}

	for input, want := range cases {
		t.Run(input, func(st *testing.T) {
			givenValue := input
			wantedValue := want
			st.Parallel()
			// When
			got := stringer.IsUpperCaseSmart(givenValue)
			// Then
			if got != wantedValue {
				st.Errorf("%q: want: %t, but got: %t", givenValue, wantedValue, got)
			}
		})
	}
}

func TestSortStringsByLength(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []string
		want  []string
	}{
		"telescopes": {
			given: []string{"Telescopes", "Glasses", "Eyes", "Monocles"},
			want:  []string{"Eyes", "Glasses", "Monocles", "Telescopes"},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// When
			got := stringer.SortByLength(testData.given)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}

func TestMiddle(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"test": {
			input: "test",
			want:  "es",
		},
		"testing": {
			input: "testing",
			want:  "t",
		},
		"middle": {
			input: "middle",
			want:  "dd",
		},
		"A": {
			input: "A",
			want:  "A",
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := stringer.GetMiddle(data.input)
			if got != data.want {
				t.Errorf("given: %s, want: %s, but got: %s", name, data.want, got)
			}
		})
	}
}

func TestScrabbleScore(t *testing.T) {
	cases := map[string]struct {
		input string
		want  int
	}{
		"empty": {
			input: "",
			want:  0,
		},
		"a": {
			input: "a",
			want:  1,
		},
		"street": {
			input: "street",
			want:  6,
		},
		"STREET": {
			input: "STREET",
			want:  6,
		},
		"st re et": {
			input: "st re et",
			want:  6,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := stringer.ScrabbleScore(data.input)
			if got != data.want {
				t.Errorf("given: %s, want: %d, but got: %d", name, data.want, got)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	cases := map[string]struct {
		input  string
		want   bool
		ending string
	}{
		"abc_c": {
			input:  "abc",
			ending: "c",
			want:   true,
		},
		"empty": {
			input:  "",
			ending: "",
			want:   true,
		},
		"space": {
			input:  " ",
			ending: "",
			want:   true,
		},
		"banana_ana": {
			input:  "banana",
			ending: "ana",
			want:   true,
		},
		"a_z": {
			input:  "a",
			ending: "z",
			want:   false,
		},
		"empty_t": {
			input:  "",
			ending: "t",
			want:   false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := stringer.EndsWith(data.input, data.ending)
			if got != data.want {
				st.Errorf("given: %s, want: %t, but got: %t", data.input, data.want, got)
			}
		})
	}
}

func TestEndsWithBrief(t *testing.T) {
	cases := map[string]struct {
		input  string
		want   bool
		ending string
	}{
		"abc_c": {
			input:  "abc",
			ending: "c",
			want:   true,
		},
		"empty": {
			input:  "",
			ending: "",
			want:   true,
		},
		"space": {
			input:  " ",
			ending: "",
			want:   true,
		},
		"banana_ana": {
			input:  "banana",
			ending: "ana",
			want:   true,
		},
		"a_z": {
			input:  "a",
			ending: "z",
			want:   false,
		},
		"empty_t": {
			input:  "",
			ending: "t",
			want:   false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := stringer.EndsWithBrief(data.input, data.ending)
			if got != data.want {
				st.Errorf("given: %s, want: %t, but got: %t", data.input, data.want, got)
			}
		})
	}
}

func TestFlickSwitch(t *testing.T) {
	// given
	cases := map[string]struct {
		list []string
		want []bool
	}{
		"flick_1": {
			list: []string{"codewars", "flick", "code", "wars"},
			want: []bool{true, false, false, false},
		},
		"flick_2": {
			list: []string{"flick", "chocolate", "adventure", "sunshine"},
			want: []bool{false, false, false, false},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// when
			got := stringer.FlickSwitch(testData.list)
			// then
			if !slices.Equal(testData.want, got) {
				st.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}
