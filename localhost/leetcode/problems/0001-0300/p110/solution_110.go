package p110

import "localhost/leetcode/util"

func isBalanced(root *util.TreeNode) bool {
	if root == nil {
		return true
	}

	left := util.BinaryTreeHeight(root.Left)
	right := util.BinaryTreeHeight(root.Right)

	ans := util.Abs(left - right)

	return ans < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}
