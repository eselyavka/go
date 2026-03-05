package p94

import "localhost/leetcode/util"

func rec_94(node *util.TreeNode, acc *[]int) {
	if node == nil {
		return
	}

	rec_94(node.Left, acc)
	*(acc) = append(*acc, node.Val)
	rec_94(node.Right, acc)
}

func inorderTraversal(root *util.TreeNode) []int {
	acc := make([]int, 0)
	rec_94(root, &acc)
	return acc
}
