package nodes_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/nodes"
)

func TestReverseKGroup(t *testing.T) {
	cases := map[string]struct {
		head *nodes.ListNode
		k    int
		want *nodes.ListNode
	}{
		"[1,2,3,4,5]_2": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: nodes.NewListNode([]int{2, 1, 4, 3, 5}),
		},
		"[1,2,3,4,5]_3": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			k:    3,
			want: nodes.NewListNode([]int{3, 2, 1, 4, 5}),
		},
		"[1,2]_2": {
			head: nodes.NewListNode([]int{1, 2}),
			k:    2,
			want: nodes.NewListNode([]int{2, 1}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := nodes.ReverseKGroup(data.head, data.k)
			compareListNodes(st, got, data.want)
		})
	}
}
