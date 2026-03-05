package p104

import "github.com/eseliavka/go/util"

func maxDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	var max int
	if left > right {
		max = left
	} else {
		max = right
	}

	return 1 + max
}
