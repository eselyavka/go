package solutions

func removeNodes(head *ListNode) *ListNode {
	ans := ListNode{Val: -1, Next: nil}
	it := &ans
	stack := make([]int, 0)

	for head != nil {
		for len(stack) > 0 && stack[len(stack)-1] < head.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}

	for _, val := range stack {
		it.Next = &ListNode{Val: val, Next: nil}
		it = it.Next
	}

	return ans.Next
}
