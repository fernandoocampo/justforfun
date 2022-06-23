package duplicates_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/duplicates"
)

func TestRemoveElement(t *testing.T) {
	cases := map[string]struct {
		nums []int
		val  int
		want int
	}{
		"2": {
			nums: []int{2},
			val:  3,
			want: 1,
		},
		"3,3": {
			nums: []int{3, 3},
			val:  5,
			want: 2,
		},
		"1,1,2,3": {
			nums: []int{1, 1, 2, 3},
			val:  1,
			want: 2,
		},
		"3, 2, 2, 3": {
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		"0,1,2,2,3,0,4,2": {
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := duplicates.RemoveElement(data.nums, data.val)
			if data.want != got {
				st.Errorf("want: %d, bug got: %d", data.want, got)
			}
		})
	}
}
