package p876

import "github.com/eseliavka/go/util"

func middleNode(head *util.ListNode) *util.ListNode {
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
