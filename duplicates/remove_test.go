package duplicates_test

import (
	"slices"
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

func TestSolve(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []int
		want  []int
	}{
		"[3, 4, 4, 3, 6, 3]": {
			given: []int{3, 4, 4, 3, 6, 3},
			want:  []int{4, 6, 3},
		},
		"[1,2,3,4]": {
			given: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// When
			got := duplicates.Solve(testData.given)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}

func TestSolveWithoutSlices(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []int
		want  []int
	}{
		"[3, 4, 4, 3, 6, 3]": {
			given: []int{3, 4, 4, 3, 6, 3},
			want:  []int{4, 6, 3},
		},
		"[1,2,3,4]": {
			given: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// When
			got := duplicates.SolveNoSlices(testData.given)
			// Then
			if !slices.Equal(testData.want, got) {
				t.Errorf("want: %+v, but got: %+v", testData.want, got)
			}
		})
	}
}
