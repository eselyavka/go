package p1448

import "github.com/eseliavka/go/util"

func rec(root *util.TreeNode, max int) int {
	if root == nil {
		return 0
	}

	var is_good int

	if root.Val >= max {
		is_good = 1
	} else {
		is_good = 0
	}

	return is_good + rec(root.Left, util.MaxInts([]int{max, root.Val})) + rec(root.Right, util.MaxInts([]int{max, root.Val}))
}

func goodNodes(root *util.TreeNode) int {
	return rec(root, root.Val)
}
