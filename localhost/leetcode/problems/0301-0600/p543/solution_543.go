package p543

import "localhost/leetcode/util"

func diameterOfBinaryTree(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	left := util.BinaryTreeHeight(root.Left)
	right := util.BinaryTreeHeight(root.Right)

	left_diameter := diameterOfBinaryTree(root.Left)
	right_diameter := diameterOfBinaryTree(root.Right)

	return util.MaxInts([]int{left + right, util.MaxInts([]int{left_diameter, right_diameter})})
}
