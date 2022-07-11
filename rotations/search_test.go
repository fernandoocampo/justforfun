package rotations_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/rotations"
)

func TestSearch(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"4,5,6,7,0,1,2:4": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := rotations.Search(data.nums, data.target)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func TestSearchOn(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"4,5,6,7,0,1,2:4": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := rotations.SearchOn(data.nums, data.target)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}
