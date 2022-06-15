package nodes

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	former := head
	head = head.Next
	return swap(head, former, nil)
}

// (1)>(2)>(3)>(4)
// (2)>(1)>(3)>(4)
// (2)>(1)>(4)>(3)
func swap(h, f, d *ListNode) *ListNode {
	if f == nil {
		return h
	}
	if f.Next == nil {
		d.Next = f
		return h
	}
	b := f.Next
	f.Next = b.Next
	b.Next = f
	f = f.Next
	if d != nil {
		d.Next = b
	}
	return swap(h, f, b.Next)
}

func SwapPairsInternet(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newhead := head.Next
	head.Next = SwapPairsInternet(newhead.Next)
	newhead.Next = head
	return newhead
}
