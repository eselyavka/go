package solutions

func rec_1448(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	var is_good int

	if root.Val >= max {
		is_good = 1
	} else {
		is_good = 0
	}

	return is_good + rec_1448(root.Left, MaxInts([]int{max, root.Val})) + rec_1448(root.Right, MaxInts([]int{max, root.Val}))
}

func goodNodes(root *TreeNode) int {
	return rec_1448(root, root.Val)
}
