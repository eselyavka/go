package p109

import "github.com/eseliavka/go/util"

func findMid(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode
	prev = nil
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	return slow
}

func sortedListToBST(head *util.ListNode) *util.TreeNode {
	if head == nil {
		return nil
	}

	mid := findMid(head)

	node := &util.TreeNode{Val: mid.Val, Left: nil, Right: nil}

	if head == mid {
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)

	return node
}
