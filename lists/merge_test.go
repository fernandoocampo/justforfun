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
		"124+134": {
			list1: lists.NewListNode([]int{1, 2, 4}),
			list2: lists.NewListNode([]int{1, 3, 4}),
			want:  lists.NewListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		"[]+[]": {
			list1: lists.NewListNode([]int{}),
			list2: lists.NewListNode([]int{}),
			want:  lists.NewListNode([]int{}),
		},
		"[]+[0]": {
			list1: lists.NewListNode([]int{}),
			list2: lists.NewListNode([]int{0}),
			want:  lists.NewListNode([]int{0}),
		},
		"[2]+[1]": {
			list1: lists.NewListNode([]int{2}),
			list2: lists.NewListNode([]int{1}),
			want:  lists.NewListNode([]int{1, 2}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := lists.MergeTwoLists(data.list1, data.list2)
			assertNodeListsEqual(st, got, data.want)
		})
	}
}

func TestMergeKLists(t *testing.T) {
	cases := map[string]struct {
		lists []*lists.ListNode
		want  *lists.ListNode
	}{
		"[[1,4,5],[1,3,4],[2,6]]": {
			lists: []*lists.ListNode{
				lists.NewListNode([]int{1, 4, 5}),
				lists.NewListNode([]int{1, 3, 4}),
				lists.NewListNode([]int{2, 6}),
			},
			want: lists.NewListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		"[]": {
			lists: []*lists.ListNode{},
			want:  lists.NewListNode([]int{}),
		},
		"[[]]": {
			lists: []*lists.ListNode{lists.NewListNode(nil)},
			want:  lists.NewListNode([]int{}),
		},
		"[[],[1]]": {
			lists: []*lists.ListNode{lists.NewListNode(nil), lists.NewListNode([]int{1})},
			want:  lists.NewListNode([]int{1}),
		},
		"[[1],[0]]": {
			lists: []*lists.ListNode{lists.NewListNode([]int{1}), lists.NewListNode([]int{0})},
			want:  lists.NewListNode([]int{0, 1}),
		},
		"[[],[-1,5,11],[],[6,10]]": {
			lists: []*lists.ListNode{
				lists.NewListNode([]int{}),
				lists.NewListNode([]int{-1, 5, 11}),
				lists.NewListNode([]int{}),
				lists.NewListNode([]int{6, 10}),
			},
			want: lists.NewListNode([]int{-1, 5, 6, 10, 11}),
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := lists.MergetKLists(data.lists)
			assertNodeListsEqual(st, got, data.want)
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

func BenchmarkMergetKLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lists.MergetKLists(
			[]*lists.ListNode{
				lists.NewListNode([]int{}),
				lists.NewListNode([]int{-1, 5, 11}),
				lists.NewListNode([]int{}),
				lists.NewListNode([]int{6, 10}),
			},
		)
	}
}

func assertNodeListsEqual(t *testing.T, l1, l2 *lists.ListNode) {
	t.Helper()
	if l1.Len() != l2.Len() {
		t.Errorf("want: %s, but l1: %s", l2, l1)
		t.FailNow()
	}
	for i, j := l1, l2; i.NextNode() != nil; i, j = i.NextNode(), j.NextNode() {
		if i.Value() != j.Value() {
			t.Errorf("want: %s, but got: %s", l2, l1)
			t.FailNow()
		}
	}
}
