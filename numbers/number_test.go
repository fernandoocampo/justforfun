package numbers_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/justforfun/numbers"
)

func TestFindSmallestNumber(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []int
		want  int
	}{
		"a": {
			given: []int{10, 8, 7, 6, 100, 13, 23},
			want:  6,
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// when
			got := numbers.FindSmallestNumber(testData.given)
			// Then
			if testData.want != got {
				st.Errorf("want: %d, but got: %d", testData.want, got)
			}
		})
	}
}

func TestGenerateFibon(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		numbers int
		want    []int
	}{
		"011235": {
			numbers: 6,
			want:    []int{0, 1, 1, 2, 3, 5},
		},
		"0112358": {
			numbers: 7,
			want:    []int{0, 1, 1, 2, 3, 5, 8},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// When
			got := numbers.MakeFibos(testData.numbers)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}
