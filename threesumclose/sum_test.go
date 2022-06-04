package threesumclose_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/threesumclose"
	"github.com/stretchr/testify/assert"
)

// expected: [][]int{[]int{1, 0, -1}, []int{1, -2, 1}, []int{0, -2, 2}, []int{0, -1, 1, 1}}
// actual mine : [][]int{[]int{-1, 0, 1}, []int{-2, 1, 1}, []int{-2, 0, 2}, []int{-1, 0, 1}}
// actual mi22 : [][]int{[]int{-2, 0, 2}, []int{-2, 1, 1}, []int{-1, 0, 1}, []int{-1, 0, 1}}
// actual copy : [][]int{[]int{0, -1, 1}, []int{1, -2, 1}, []int{0, -2, 2}}

func TestThreeSum(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"-1,2,1,-4": {
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		"0,0,0": {
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		"1,1,1,0": {
			nums:   []int{1, 1, 1, 0},
			target: -100,
			want:   2,
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := threesumclose.ThreeSum(data.nums, data.target)
			assert.Equal(st, data.want, got)
		})
	}
}
