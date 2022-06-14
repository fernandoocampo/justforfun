package lists

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

func MergetKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeLists(lists, nil)
}

func mergeLists(lists []*ListNode, out *ListNode) *ListNode {
	if len(lists) == 0 {
		return out
	}

	var l2 []*ListNode
	for _, node := range lists {
		if node == nil {
			continue
		}
		out = feedNode(node.Val, out, out)
		l2 = append(l2, node.Next)
	}

	return mergeLists(l2, out)
}

func feedNode(val int, head, former *ListNode) *ListNode {
	if former == nil {
		return &ListNode{Val: val}
	}
	if former != nil && val < former.Val {
		newNode := &ListNode{
			Val:  val,
			Next: former,
		}
		if head == former {
			former = newNode
			return former
		}
		return head
	}
	if former != nil && former.Next == nil {
		former.Next = &ListNode{Val: val}
		return head
	}
	if val < former.Next.Val {
		newNode := &ListNode{
			Val:  val,
			Next: former.Next,
		}
		former.Next = newNode
		return head
	}
	return feedNode(val, head, former.Next)
}
