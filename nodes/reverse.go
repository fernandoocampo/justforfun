package nodes

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	currhead := head
	newhead, ok := tail(currhead, k)
	if !ok {
		return currhead
	}
	newGroupHead := newhead.Next
	reversedNode := reverse(head, k)
	reversedNode.Next = ReverseKGroup(newGroupHead, k)
	return newhead
}

func tail(head *ListNode, k int) (*ListNode, bool) {
	if head == nil {
		return head, false
	}
	if k == 1 {
		return head, true
	}
	return tail(head.Next, k-1)
}

func reverse(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	newhead := reverse(head.Next, k-1)
	newhead.Next = head
	return head
}
