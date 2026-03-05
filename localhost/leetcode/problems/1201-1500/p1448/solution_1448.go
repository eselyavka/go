package p1448

import "localhost/leetcode/util"

func rec_1448(root *util.TreeNode, max int) int {
	if root == nil {
		return 0
	}

	var is_good int

	if root.Val >= max {
		is_good = 1
	} else {
		is_good = 0
	}

	return is_good + rec_1448(root.Left, util.MaxInts([]int{max, root.Val})) + rec_1448(root.Right, util.MaxInts([]int{max, root.Val}))
}

func goodNodes(root *util.TreeNode) int {
	return rec_1448(root, root.Val)
}
