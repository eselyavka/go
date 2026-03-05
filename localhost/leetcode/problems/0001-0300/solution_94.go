package solutions

func rec_94(node *TreeNode, acc *[]int) {
	if node == nil {
		return
	}

	rec_94(node.Left, acc)
	*(acc) = append(*acc, node.Val)
	rec_94(node.Right, acc)
}

func inorderTraversal(root *TreeNode) []int {
	acc := make([]int, 0)
	rec_94(root, &acc)
	return acc
}
