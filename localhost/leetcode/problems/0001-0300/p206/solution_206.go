package p206

import "localhost/leetcode/util"

func reverseList(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode
	curr := head
	var next *util.ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
