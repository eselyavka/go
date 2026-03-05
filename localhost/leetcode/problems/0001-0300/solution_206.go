package solutions

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
