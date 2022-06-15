package nodes_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/nodes"
)

func TestSwapPairs(t *testing.T) {
	cases := map[string]struct {
		head *nodes.ListNode
		want *nodes.ListNode
	}{
		"[1,2,3,4]": {
			head: nodes.NewListNode([]int{1, 2, 3, 4}),
			want: nodes.NewListNode([]int{2, 1, 4, 3}),
		},
		"[1,2,3,4,5]": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			want: nodes.NewListNode([]int{2, 1, 4, 3, 5}),
		},
		"[1,2,3]": {
			head: nodes.NewListNode([]int{1, 2, 3}),
			want: nodes.NewListNode([]int{2, 1, 3}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := nodes.SwapPairs(data.head)
			compareListNodes(st, got, data.want)
		})
	}
}

func TestSwapPairsInternet(t *testing.T) {
	cases := map[string]struct {
		head *nodes.ListNode
		want *nodes.ListNode
	}{
		"[1,2,3,4]": {
			head: nodes.NewListNode([]int{1, 2, 3, 4}),
			want: nodes.NewListNode([]int{2, 1, 4, 3}),
		},
		"[1,2,3,4,5]": {
			head: nodes.NewListNode([]int{1, 2, 3, 4, 5}),
			want: nodes.NewListNode([]int{2, 1, 4, 3, 5}),
		},
		"[1,2,3]": {
			head: nodes.NewListNode([]int{1, 2, 3}),
			want: nodes.NewListNode([]int{2, 1, 3}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := nodes.SwapPairsInternet(data.head)
			compareListNodes(st, got, data.want)
		})
	}
}

func compareListNodes(t *testing.T, got, want *nodes.ListNode) {
	t.Helper()
	if got.Len() != want.Len() {
		t.Errorf("want: %s, but got: %s", want, got)
		t.FailNow()
	}
	for i, j := want, got; i.NextNode() != nil; i, j = i.NextNode(), j.NextNode() {
		if i.Value() != j.Value() {
			t.Errorf("want: %s, but got: %s", want, got)
			t.FailNow()
		}
	}
}

func BenchmarkSwapPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := nodes.NewListNode([]int{1, 2, 3, 4, 5})
		_ = nodes.SwapPairs(head)
		head = nil
	}
}

func BenchmarkSwapPairsInternet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := nodes.NewListNode([]int{1, 2, 3, 4, 5})
		_ = nodes.SwapPairsInternet(head)
		head = nil
	}
}
