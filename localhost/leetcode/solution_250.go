package solutions

func rec_250(node *TreeNode, cnt *int) bool {
	if node == nil {
		return true
	}

	left := rec_250(node.Left, cnt)
	right := rec_250(node.Right, cnt)

	if !left || !right {
		return false
	}

	if node.Left != nil && node.Left.Val != node.Val {
		return false
	}

	if node.Right != nil && node.Right.Val != node.Val {
		return false
	}

	*cnt++

	return true

}
func countUnivalSubtrees(root *TreeNode) int {
	cnt := 0
	rec_250(root, &cnt)

	return cnt
}
