package p141

import "localhost/leetcode/util"

func hasCycle(head *util.ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
