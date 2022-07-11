package finders_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/finders"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"empty": {
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		"1_1": {
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		"5,7,7,8,8,10_6": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		"5,7,7,8,8,10_8": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		"5,7,7,8,8,8,10_8": {
			nums:   []int{5, 7, 7, 8, 8, 8, 10},
			target: 8,
			want:   []int{3, 5},
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := finders.SearchRange(data.nums, data.target)
			assert.Equal(st, data.want, got)
		})
	}
}

func TestSearchRangeOn(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"empty": {
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		"1_1": {
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		"5,7,7,8,8,10_6": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		"5,7,7,8,8,10_8": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		"5,7,7,8,8,8,10_8": {
			nums:   []int{5, 7, 7, 8, 8, 8, 10},
			target: 8,
			want:   []int{3, 5},
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := finders.SearchRangeOn(data.nums, data.target)
			assert.Equal(st, data.want, got)
		})
	}
}
