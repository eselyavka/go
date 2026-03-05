package p94

import "localhost/leetcode/util"

func rec(node *util.TreeNode, acc *[]int) {
	if node == nil {
		return
	}

	rec(node.Left, acc)
	*(acc) = append(*acc, node.Val)
	rec(node.Right, acc)
}

func inorderTraversal(root *util.TreeNode) []int {
	acc := make([]int, 0)
	rec(root, &acc)
	return acc
}
