package p708

import "github.com/eseliavka/go/util"

func insert(aNode *util.ListNode, x int) *util.ListNode {
	if aNode == nil {
		node := util.ListNode{Val: x, Next: nil}
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
			prev.Next = &util.ListNode{Val: x, Next: curr}
			return aNode
		}

		prev = curr
		curr = curr.Next
	}

	prev.Next = &util.ListNode{Val: x, Next: curr}

	return aNode
}
