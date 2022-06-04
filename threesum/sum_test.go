package threesum_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/threesum"
	"github.com/stretchr/testify/assert"
)

// expected: [][]int{[]int{1, 0, -1}, []int{1, -2, 1}, []int{0, -2, 2}, []int{0, -1, 1, 1}}
// actual mine : [][]int{[]int{-1, 0, 1}, []int{-2, 1, 1}, []int{-2, 0, 2}, []int{-1, 0, 1}}
// actual mi22 : [][]int{[]int{-2, 0, 2}, []int{-2, 1, 1}, []int{-1, 0, 1}, []int{-1, 0, 1}}
// actual copy : [][]int{[]int{0, -1, 1}, []int{1, -2, 1}, []int{0, -2, 2}}

func TestThreeSum(t *testing.T) {
	cases := map[string]struct {
		nums []int
		want [][]int
	}{
		"1,0,-2,-1,1,2": {
			nums: []int{1, 0, -2, -1, 1, 2},
			want: [][]int{{1, 0, -1}, {1, -2, 1}, {0, -2, 2}, {0, -1, 1, 1}},
		},
		// "-1,0,1,2,-1,-4": {
		// 	nums: []int{-1, 0, 1, 2, -1, -4},
		// 	want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		// },
		// "empty": {
		// 	nums: []int{},
		// 	want: [][]int{},
		// },
		// "0": {
		// 	nums: []int{0},
		// 	want: [][]int{},
		// },
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := threesum.ThreeSum(data.nums)
			assert.Equal(st, data.want, got)
		})
	}
}
