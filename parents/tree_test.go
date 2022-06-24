package parents_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/parents"
	"github.com/stretchr/testify/assert"
)

func TestGetNode(t *testing.T) {
	cases := map[string]struct {
		root *parents.TreeNode
		want int
	}{
		"5": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}),
			want: 5,
		},
		"11": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}),
			want: 11,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parents.GetNode(data.root, data.want)
			if data.want != got.Key {
				st.Errorf("want: %d, but got: %d", data.want, got.Key)
			}
		})
	}
}

func TestGetNodes(t *testing.T) {
	cases := map[string]struct {
		root *parents.TreeNode
		want []int
	}{
		"5,11": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}),
			want: []int{5, 11},
		},
		"11,-3": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}),
			want: []int{-3, 2},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parents.GetNodes(data.root, data.want...)
			assert.Equal(st, data.want, got)
		})
	}
}
