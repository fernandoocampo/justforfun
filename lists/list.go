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
