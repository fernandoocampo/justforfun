package lists_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/lists"
)

func TestMergeTwoLists(t *testing.T) {
	cases := map[string]struct {
		list1 *lists.ListNode
		list2 *lists.ListNode
		want  *lists.ListNode
	}{
		// "124+134": {
		// 	list1: lists.NewListNode([]int{1, 2, 4}),
		// 	list2: lists.NewListNode([]int{1, 3, 4}),
		// 	want:  lists.NewListNode([]int{1, 1, 2, 3, 4, 4}),
		// },
		// "[]+[]": {
		// 	list1: lists.NewListNode([]int{}),
		// 	list2: lists.NewListNode([]int{}),
		// 	want:  lists.NewListNode([]int{}),
		// },
		// "[]+[0]": {
		// 	list1: lists.NewListNode([]int{}),
		// 	list2: lists.NewListNode([]int{0}),
		// 	want:  lists.NewListNode([]int{0}),
		// },
		"[2]+[1]": {
			list1: lists.NewListNode([]int{2}),
			list2: lists.NewListNode([]int{1}),
			want:  lists.NewListNode([]int{1, 2}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := lists.MergeTwoLists(data.list1, data.list2)
			if got.Len() != data.want.Len() {
				st.Errorf("want: %s, but got: %s", data.want, got)
				st.FailNow()
			}
			for i, j := got, data.want; i.NextNode() != nil; i, j = i.NextNode(), j.NextNode() {
				if i.Value() != j.Value() {
					st.Errorf("want: %s, but got: %s", data.want, got)
					st.FailNow()
				}
			}
		})
	}
}

func BenchmarkMergetTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lists.MergeTwoLists(
			lists.NewListNode([]int{1, 2, 4}),
			lists.NewListNode([]int{1, 3, 4}),
		)
	}
}
