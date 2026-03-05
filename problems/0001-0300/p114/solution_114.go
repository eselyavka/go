package p114

import "github.com/eseliavka/go/util"

func flatten(root *util.TreeNode) {
	if root == nil {
		return
	}

	node := root
	for node != nil {
		if node.Left != nil {
			rightmost := node.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}

			rightmost.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}

		node = node.Right
	}
}
