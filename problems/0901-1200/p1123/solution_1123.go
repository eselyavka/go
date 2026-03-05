package p1123

import "github.com/eseliavka/go/util"

func rec(node *util.TreeNode, curr, height int) *util.TreeNode {
	if node == nil || curr == height {
		return node
	}
	left := rec(node.Left, curr+1, height)
	right := rec(node.Right, curr+1, height)

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

	return rec(root, 1, height)
}
