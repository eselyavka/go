package solutions

func swapNodes(head *ListNode, k int) *ListNode {
	length := 0

	it := head

	var left_node *ListNode = nil
	for it != nil {
		length++
		if length == k {
			left_node = it
		}
		it = it.Next
	}

	it = head

	var right_node *ListNode = nil
	for i := 0; i <= length-k; i++ {
		right_node = it
		it = it.Next
	}

	tmp := left_node.Val
	left_node.Val = right_node.Val
	right_node.Val = tmp

	return head
}
