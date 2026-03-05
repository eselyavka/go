package p2

import "localhost/leetcode/util"

// addTwoNumbers adds two non-negative integers represented in reverse-order lists.
func addTwoNumbers(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &util.ListNode{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}

	return dummy.Next
}
