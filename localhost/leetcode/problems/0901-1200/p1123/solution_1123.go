package p1123

import "localhost/leetcode/util"

func rec_1123(node *util.TreeNode, curr, height int) *util.TreeNode {
	if node == nil || curr == height {
		return node
	}
	left := rec_1123(node.Left, curr+1, height)
	right := rec_1123(node.Right, curr+1, height)

	if left != nil && right != nil {
		return node
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}
func lcaDeepestLeaves(root *util.TreeNode) *util.TreeNode {
	height := util.BinaryTreeHeight(root)

	return rec_1123(root, 1, height)
}
