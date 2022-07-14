package positions_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/positions"
)

func TestSearchInsert(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"1,3,5,6:5": {
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		"1,3,5,6:2": {
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		"1,3,5,6:7": {
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		"0,1,3,4,5,6,7,8,9,10,11,12,13:2": {
			nums:   []int{0, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			target: 2,
			want:   2,
		},
		"1:0": {
			nums:   []int{1},
			target: 0,
			want:   0,
		},
		"1:2": {
			nums:   []int{1},
			target: 2,
			want:   1,
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := positions.SearchInsert(data.nums, data.target)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}
