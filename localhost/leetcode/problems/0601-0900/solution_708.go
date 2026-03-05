package solutions

func insert_708(aNode *ListNode, x int) *ListNode {
	if aNode == nil {
		node := ListNode{Val: x, Next: nil}
		node.Next = &node

		return &node
	}

	curr := aNode.Next
	prev := aNode

	done := false

	for curr != aNode {
		if x >= prev.Val && x <= curr.Val {
			done = true
		} else if prev.Val > curr.Val {
			if x >= prev.Val || x <= curr.Val {
				done = true
			}
		}

		if done {
			prev.Next = &ListNode{Val: x, Next: curr}
			return aNode
		}

		prev = curr
		curr = curr.Next
	}

	prev.Next = &ListNode{Val: x, Next: curr}

	return aNode
}
