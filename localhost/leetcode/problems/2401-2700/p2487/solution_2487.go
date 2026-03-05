package p2487

import "localhost/leetcode/util"

func removeNodes(head *util.ListNode) *util.ListNode {
	ans := util.ListNode{Val: -1, Next: nil}
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
		it.Next = &util.ListNode{Val: val, Next: nil}
		it = it.Next
	}

	return ans.Next
}
