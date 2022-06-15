package nodes

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	pen, nodes := head.penultimate()
	if n == 1 {
		pen.Next = nil
		return head
	}
	if n == nodes {
		head = head.Next
		return head
	}
	before := head
	after := head.Next.Next
	for i := nodes; i != n+1; i-- {
		before = before.Next
		if before.Next == nil {
			after = nil
			break
		}
		after = before.Next.Next
	}
	before.Next = after
	return head
}

func (l *ListNode) penultimate() (*ListNode, int) {
	if l == nil {
		return nil, 0
	}
	var result *ListNode
	count := 1
	for i := l; i.Next != nil; i = i.Next {
		result = i
		count++
	}
	return result, count
}
