package p206

import "github.com/eseliavka/go/util"

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
