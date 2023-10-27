package numbers_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/numbers"
)

func TestFindSmallestNumber(t *testing.T) {
	t.Parallel()
	// Given
	cases := map[string]struct {
		given []int
		want  int
	}{
		"a": {
			given: []int{10, 8, 7, 6, 100, 13, 23},
			want:  6,
		},
	}

	for testName, testData := range cases {
		t.Run(testName, func(st *testing.T) {
			testData := testData
			st.Parallel()
			// when
			got := numbers.FindSmallestNumber(testData.given)
			// Then
			if testData.want != got {
				st.Errorf("want: %d, but got: %d", testData.want, got)
			}
		})
	}

}
