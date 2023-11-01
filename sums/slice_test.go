package sums_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/justforfun/sums"
)

func TestBetween(t *testing.T) {
	// Given
	cases := map[string]struct {
		start int
		end   int
		want  []int
	}{
		"1_4": {
			start: 1,
			end:   4,
			want:  []int{1, 2, 3, 4},
		},
		"-2_2": {
			start: -2,
			end:   2,
			want:  []int{-2, -1, 0, 1, 2},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// When
			got := sums.Between(testData.start, testData.end)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}

func TestSquareSum(t *testing.T) {
	// Given
	cases := map[string]struct {
		numbers []int
		want    int
	}{
		"[1,2]": {
			numbers: []int{1, 2},
			want:    5,
		},
		"[0,3,4,5]": {
			numbers: []int{0, 3, 4, 5},
			want:    50,
		},
		"[]": {
			numbers: []int{},
			want:    0,
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			// When
			got := sums.SquareSum(testData.numbers)
			// Then
			if testData.want != got {
				st.Errorf("want: %d, but got: %d", testData.want, got)
			}
		})
	}
}
