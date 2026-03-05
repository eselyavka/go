package p1721

import "localhost/leetcode/util"

func swapNodes(head *util.ListNode, k int) *util.ListNode {
	length := 0

	it := head

	var left_node *util.ListNode = nil
	for it != nil {
		length++
		if length == k {
			left_node = it
		}
		it = it.Next
	}

	it = head

	var right_node *util.ListNode = nil
	for i := 0; i <= length-k; i++ {
		right_node = it
		it = it.Next
	}

	tmp := left_node.Val
	left_node.Val = right_node.Val
	right_node.Val = tmp

	return head
}
