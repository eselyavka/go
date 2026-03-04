package solutions

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := 0
	it := head
	for it != nil {
		sz++
		it = it.Next
	}

	curr := head
	var prev *ListNode
	var next *ListNode

	i := 0
	for curr != nil {
		if sz-n == 0 {
			return curr.Next
		}

		if sz-n == i {
			next = curr.Next
			prev.Next = next
			break
		}
		prev = curr
		curr = curr.Next
		i++
	}

	return head
}

func removeNthFromEndFast(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	left := dummy
	right := head

	for n > 0 {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}
