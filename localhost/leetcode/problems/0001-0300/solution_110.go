package solutions

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := binaryTreeHeight(root.Left)
	right := binaryTreeHeight(root.Right)

	ans := Abs(left - right)

	return ans < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}
