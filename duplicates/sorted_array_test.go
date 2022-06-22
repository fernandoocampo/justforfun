package duplicates_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/duplicates"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := map[string]struct {
		nums []int
		want int
	}{
		"1,1,2": {
			nums: []int{1, 1, 2},
			want: 2,
		},
		"0,0,1,1,1,2,2,3,3,4": {
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		"1,1": {
			nums: []int{1, 1},
			want: 1,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := duplicates.RemoveDuplicates(data.nums)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}

func TestRemoveDuplicatesInternet(t *testing.T) {
	cases := map[string]struct {
		nums []int
		want int
	}{
		"1,1,2": {
			nums: []int{1, 1, 2},
			want: 2,
		},
		"0,0,1,1,1,2,2,3,3,4": {
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		"1,1": {
			nums: []int{1, 1},
			want: 1,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := duplicates.RemoveDuplicatesInternet(data.nums)
			if got != data.want {
				st.Errorf("want: %d, but got: %d", data.want, got)
			}
		})
	}
}
