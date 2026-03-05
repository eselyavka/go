package solutions

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	arr := make([]*ListNode, 0)
	it := head
	for it != nil {
		next := it.Next
		it.Next = nil
		arr = append(arr, it)

		it = next
	}

	for i := 1; i < len(arr); i += 2 {
		tmp := arr[i]
		arr[i] = arr[i-1]
		arr[i-1] = tmp
	}

	new_head := &ListNode{Val: -1, Next: nil}
	tail := new_head

	for _, node := range arr {
		tail.Next = node
		tail = tail.Next
	}

	return new_head.Next
}
