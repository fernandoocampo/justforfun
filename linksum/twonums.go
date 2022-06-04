package linksum

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	root := ListNode{
		Value: values[0],
	}
	node := &root
	for i := 1; i < len(values); i++ {
		newNode := ListNode{
			Value: values[i],
		}
		node.Next = &newNode
		node = &newNode
	}
	return &root
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	var builder strings.Builder
	node := l
	for node != nil {
		builder.WriteString(strconv.Itoa(node.Value))
		node = node.Next
	}
	return builder.String()
}

func AddTwoNumbers1(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l1Int := getNumberFromList(l1, 1)
	l2Int := getNumberFromList(l2, 1)
	sum := l1Int + l2Int
	if sum == 0 {
		return &ListNode{}
	}
	return getListNodeFromInt(sum)
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := sumTwoList(l1, l2, 0)
	return result
}

func sumTwoList(l1, l2 *ListNode, remain int) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var l1val, l2val int
	var l1next, l2next *ListNode
	if l1 != nil {
		l1val = l1.Value
		l1next = l1.Next
	}
	if l2 != nil {
		l2val = l2.Value
		l2next = l2.Next
	}
	sum := l1val + l2val + remain
	remain = 0
	if sum >= 10 {
		sum = sum % 10
		remain = 1
	}
	nextNode := sumTwoList(l1next, l2next, remain)
	if nextNode == nil && remain > 0 {
		nextNode = &ListNode{
			Value: remain,
		}
	}
	node := ListNode{
		Value: sum,
		Next:  nextNode,
	}
	return &node
}

func getNumberFromList(l *ListNode, base int) int {
	if l == nil {
		return 0
	}
	return l.Value*base + getNumberFromList(l.Next, base*10)
}

func getListNodeFromInt(value int) *ListNode {
	if value == 0 {
		return nil
	}
	node := ListNode{
		Value: value % 10,
		Next:  getListNodeFromInt(value / 10),
	}
	return &node
}
