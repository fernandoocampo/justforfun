package twosum_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/twosum"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"0_1": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		"1_2": {
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		"3_3": {
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		"empty_1": {
			nums:   []int{1, 3},
			target: 6,
			want:   []int{},
		},
		"empty": {
			nums:   []int{},
			target: 6,
			want:   []int{},
		},
	}

	for name, testdata := range cases {
		givenData := testdata
		t.Run(name, func(st *testing.T) {
			st.Parallel()

			got := twosum.TwoSum(givenData.nums, givenData.target)
			if len(got) != len(givenData.want) {
				st.Errorf("want: %+v, but got: %+v", givenData.want, got)
				st.FailNow()
			}
			for idx, v := range givenData.want {
				if v != got[idx] {
					st.Errorf("want: %+v, but got: %+v", givenData.want, got)
				}
			}
		})
	}
}
