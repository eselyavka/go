package solutions

func invert(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		return
	}

	left := node.Left
	right := node.Right

	node.Left = right
	node.Right = left

	invert(node.Left)
	invert(node.Right)

	return
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)

	return root
}
