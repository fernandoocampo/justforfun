package lists

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values []int) *ListNode {
	var head *ListNode
	if len(values) == 0 {
		return head
	}
	head = &ListNode{}
	head.Val = values[0]
	follower := head
	for i := 1; i < len(values); i++ {
		newnode := ListNode{
			Val: values[i],
		}
		follower.Next = &newnode
		follower = &newnode
	}
	return head
}

func (l *ListNode) Value() int {
	if l != nil {
		return l.Val
	}
	return 0
}

func (l *ListNode) NextNode() *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func (l *ListNode) Len() int {
	if l == nil {
		return 0
	}
	count := 0
	for i := l; ; i = i.Next {
		if i == nil {
			break
		}
		count++
	}
	return count
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	b := strings.Builder{}
	for i := l; ; i = i.NextNode() {
		if i == nil {
			break
		}
		b.WriteString(strconv.Itoa(i.Value()))
		if i.NextNode() != nil {
			b.WriteRune('-')
		}
	}
	return b.String()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var result *ListNode
	switch {
	case list1.Val > list2.Val:
		result = &ListNode{Val: list2.Val}
		list2 = list2.Next
	case list2.Val > list1.Val:
		result = &ListNode{Val: list1.Val}
		list1 = list1.Next
	default:
		result = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}
	newhead := result
	merge(list1, list2, newhead)
	return result
}

func merge(list1, list2, new *ListNode) {
	if list1 == nil && list2 == nil {
		return
	}
	if list1 == nil {
		new.Next = list2
		merge(list1, list2.Next, new.Next)
		return
	}
	if list2 == nil {
		new.Next = &ListNode{Val: list1.Val}
		merge(list1.Next, list2, new.Next)
		return
	}
	if list1.Val > list2.Val {
		new.Next = &ListNode{Val: list2.Val}
		merge(list1, list2.Next, new.Next)
		return
	}
	if list2.Val > list1.Val {
		new.Next = &ListNode{Val: list1.Val}
		merge(list1.Next, list2, new.Next)
		return
	}
	new.Next = &ListNode{Val: list1.Val}
	new = new.Next
	new.Next = &ListNode{Val: list2.Val}
	merge(list1.Next, list2.Next, new.Next)
}
