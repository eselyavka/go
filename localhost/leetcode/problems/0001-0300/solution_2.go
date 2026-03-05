package solutions

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{Val: -1, Next: nil}
	ans := l3
	rem := 0

	for l1 != nil || l2 != nil {
		sum_1 := 0
		sum_2 := 0

		if l1 != nil {
			sum_1 = l1.Val
		}

		if l2 != nil {
			sum_2 = l2.Val
		}

		sum := sum_1 + sum_2 + rem

		if rem == 1 {
			rem = 0
		}

		if sum >= 10 {
			sum %= 10
			rem = 1
		}

		for l3.Next != nil {
			l3 = l3.Next
		}

		l3.Next = &ListNode{Val: sum, Next: nil}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if rem == 1 {
		l3.Next.Next = &ListNode{Val: rem, Next: nil}
	}

	return ans.Next
}
