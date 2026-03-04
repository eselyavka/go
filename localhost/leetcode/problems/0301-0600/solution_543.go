package solutions

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := binaryTreeHeight(root.Left)
	right := binaryTreeHeight(root.Right)

	left_diameter := diameterOfBinaryTree(root.Left)
	right_diameter := diameterOfBinaryTree(root.Right)

	return MaxInts([]int{left + right, MaxInts([]int{left_diameter, right_diameter})})
}
