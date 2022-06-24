package parents_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/parents"
)

func TestGetCommonParent(t *testing.T) {
	cases := map[string]struct {
		root *parents.TreeNode
		cone int
		ctwo int
		want int
	}{
		"3_1=5": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 13, 2, 6, 11, 3, -2, 7, 1}),
			cone: 3,
			ctwo: 1,
			want: 5,
		},
		"5_-3=10": {
			root: parents.NewTreeWithLeafs([]int{10, 5, -3, 13, 2, 6, 11, 3, -2, 7, 1}),
			cone: 5,
			ctwo: -3,
			want: 10,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			nodes := parents.GetNodes(data.root, data.cone, data.ctwo, data.want)
			got := parents.GetCommonParent(nodes[0], nodes[1])
			if nodes[2] != got {
				st.Errorf("want: %d, but got: %v", data.want, got)
			}
		})
	}
}
