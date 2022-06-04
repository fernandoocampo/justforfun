package nodes_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/nodes"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := map[string]struct {
		head *nodes.ListNode
		n    int
		want *nodes.ListNode
	}{
		"[1,2,3,4,5]_1": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			n:    1,
			want: nodes.NewListNode([]int{1, 2, 3, 4}),
		},
		"[1,2,3,4,5]_2": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			n:    2,
			want: nodes.NewListNode([]int{1, 2, 3, 5}),
		},
		"[1]": {
			head: nodes.NewListNode([]int{1}),
			n:    1,
			want: nil,
		},
		"[1,2]_1": {
			head: nodes.NewListNode([]int{1, 2}),
			n:    1,
			want: nodes.NewListNode([]int{1}),
		},
		"[1,2]_2": {
			head: nodes.NewListNode([]int{1, 2}),
			n:    2,
			want: nodes.NewListNode([]int{2}),
		},
	}
	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := nodes.RemoveNthFromEnd(data.head, data.n)
			if got.Len() != data.want.Len() {
				st.Errorf("want: %s, but got: %s", data.want, got)
				st.FailNow()
			}
			for i, j := data.want, got; i.NextNode() != nil; i, j = i.NextNode(), j.NextNode() {
				if i.Value() != j.Value() {
					st.Errorf("want: %s, but got: %s", data.want, got)
					st.FailNow()
				}
			}
		})
	}
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node := nodes.NewListNode([]int{1, 2, 3, 4, 5})
		nodes.RemoveNthFromEnd(node, 2)
	}
}
