package solutions

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: -101, Next: nil}

	tail := head
	var next *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next = &ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			next = &ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}
		tail.Next = next
		tail = tail.Next

	}

	for list1 != nil {
		next := &ListNode{Val: list1.Val, Next: nil}
		tail.Next = next
		tail = tail.Next
		list1 = list1.Next
	}

	for list2 != nil {
		next := &ListNode{Val: list2.Val, Next: nil}
		tail.Next = next
		tail = tail.Next
		list2 = list2.Next
	}

	return head.Next
}
