package p250

import "localhost/leetcode/util"

func rec(node *util.TreeNode, cnt *int) bool {
	if node == nil {
		return true
	}

	left := rec(node.Left, cnt)
	right := rec(node.Right, cnt)

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
func countUnivalSubtrees(root *util.TreeNode) int {
	cnt := 0
	rec(root, &cnt)

	return cnt
}
