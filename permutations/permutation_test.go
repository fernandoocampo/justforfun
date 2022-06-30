package permutations_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/permutations"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	cases := map[string]struct {
		nums []int
		want []int
	}{
		"empty": {
			nums: nil,
			want: nil,
		},
		"1": {
			nums: []int{1},
			want: []int{1},
		},
		"1,2": {
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		"1,2,3": {
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		"1,3,2": {
			nums: []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
		"3,2,1": {
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		"1,1,5": {
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		"1,5,2,5": {
			nums: []int{1, 5, 2, 5},
			want: []int{1, 5, 5, 2},
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			permutations.NextPermutation(data.nums)
			assert.Equal(st, data.want, data.nums)
		})
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = []int{2, 4, 5, 1, 7, 6, 3}
		permutations.NextPermutation(result)
	}
	b.Log(result)
}
