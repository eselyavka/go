package p21

import "github.com/eseliavka/go/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	head := &util.ListNode{Val: -101, Next: nil}

	tail := head
	var next *util.ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next = &util.ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			next = &util.ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}
		tail.Next = next
		tail = tail.Next

	}

	for list1 != nil {
		next := &util.ListNode{Val: list1.Val, Next: nil}
		tail.Next = next
		tail = tail.Next
		list1 = list1.Next
	}

	for list2 != nil {
		next := &util.ListNode{Val: list2.Val, Next: nil}
		tail.Next = next
		tail = tail.Next
		list2 = list2.Next
	}

	return head.Next
}
