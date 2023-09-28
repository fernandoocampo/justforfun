package linksum_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/linksum"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		l1   *linksum.ListNode
		l2   *linksum.ListNode
		want *linksum.ListNode
	}{
		"7_0_8": {
			l1:   linksum.NewListNode([]int{2, 4, 3}),
			l2:   linksum.NewListNode([]int{5, 6, 4}),
			want: linksum.NewListNode([]int{7, 0, 8}),
		},
		"7_0_0": {
			l1:   linksum.NewListNode([]int{0, 0, 5}),
			l2:   linksum.NewListNode([]int{0, 0, 2}),
			want: linksum.NewListNode([]int{0, 0, 7}),
		},
		"0": {
			l1:   linksum.NewListNode([]int{0}),
			l2:   linksum.NewListNode([]int{0}),
			want: linksum.NewListNode([]int{0}),
		},
		"8,9,9,9,0,0,0,1": {
			l1:   linksum.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   linksum.NewListNode([]int{9, 9, 9, 9}),
			want: linksum.NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		"6,6,4,..0,...,1": {
			l1:   linksum.NewListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2:   linksum.NewListNode([]int{5, 6, 4}),
			want: linksum.NewListNode([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		"1,0,0,0,1": {
			l1:   linksum.NewListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2:   linksum.NewListNode([]int{5, 6, 4}),
			want: linksum.NewListNode([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	for name, data := range cases {
		testdata := data

		t.Run(name, func(st *testing.T) {
			st.Parallel()

			got := linksum.AddTwoNumbers(testdata.l1, testdata.l2)
			compareListNodes(st, got, testdata.want)
		})
	}
}

func compareListNodes(t *testing.T, got, want *linksum.ListNode) {
	t.Helper()
	t.Logf("receiving want: %+v, got: %+v", want, got)
	gotCount := countListNodes(got)
	wantCount := countListNodes(want)
	if gotCount != wantCount {
		t.Errorf("counting want: %s, but got: %s", want, got)
		t.FailNow()
	}
	wantNode := want
	gotNode := got
	for wantNode != nil && gotNode != nil {
		if wantNode.Value != gotNode.Value {
			t.Logf("failed with want: %+v, got: %+v", wantNode, gotNode)
			t.Errorf("comparing want: %s, but got: %s", want, got)
			t.FailNow()
		}
		wantNode = wantNode.Next
		gotNode = gotNode.Next
	}
}

func countListNodes(list *linksum.ListNode) int {
	var count int
	if list == nil {
		return count
	}
	node := list
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
