package p129

import "github.com/eseliavka/go/util"

func sumNumbers(root *util.TreeNode) int {
	var dfs func(node *util.TreeNode, curr int) int
	dfs = func(node *util.TreeNode, curr int) int {
		if node == nil {
			return 0
		}

		curr = curr*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return curr
		}

		return dfs(node.Left, curr) + dfs(node.Right, curr)
	}

	return dfs(root, 0)
}
