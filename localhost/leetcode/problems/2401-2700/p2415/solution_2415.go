package p2415

import "localhost/leetcode/util"

func invertOdd(root1, root2 *util.TreeNode, level int) {
	if root1 == nil || root2 == nil {
		return
	}

	if level%2 == 0 {
		temp := root2.Val
		root2.Val = root1.Val
		root1.Val = temp
	}

	invertOdd(root1.Left, root2.Right, level+1)
	invertOdd(root1.Right, root2.Left, level+1)
}

func reverseOddLevels(root *util.TreeNode) *util.TreeNode {
	if root.Left != nil && root.Right != nil {
		invertOdd(root.Left, root.Right, 0)
	}
	return root
}
