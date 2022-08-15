package combination_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/combination"
	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	cases := map[string]struct {
		candidates []int
		target     int
		want       [][]int
	}{
		"2,3,6,7-7": {
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{3, 2, 2}, {7}},
		},
		"2,3,5-8": {
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		"1,2-3": {
			candidates: []int{1, 2},
			target:     3,
			want:       [][]int{{1, 1, 1}, {1, 2}},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := combination.CombinationSum(data.candidates, data.target)
			assert.Equal(st, data.want, got)
		})
	}
}

func TestCombinationSumInternet(t *testing.T) {
	cases := map[string]struct {
		candidates []int
		target     int
		want       [][]int
	}{
		"2,3,6,7-7": {
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		"2,3,5-8": {
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		"1,2-3": {
			candidates: []int{1, 2},
			target:     3,
			want:       [][]int{{1, 1, 1}, {1, 2}},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := combination.CombinationSumInternet(data.candidates, data.target)
			assert.Equal(st, data.want, got)
		})
	}
}
