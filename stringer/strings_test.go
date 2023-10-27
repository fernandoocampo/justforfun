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
