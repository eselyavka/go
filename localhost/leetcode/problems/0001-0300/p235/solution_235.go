package p235

import "localhost/leetcode/util"

func dfs_235(root, node *util.TreeNode, path *[]*util.TreeNode) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root)

	if root.Val == node.Val {
		return true
	}

	if dfs_235(root.Left, node, path) || dfs_235(root.Right, node, path) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func lowestCommonAncestorTreeNode(root, p, q *util.TreeNode) *util.TreeNode {
	p_to_p := make([]*util.TreeNode, 0)
	p_to_q := make([]*util.TreeNode, 0)
	_ = dfs_235(root, p, &p_to_p)
	_ = dfs_235(root, q, &p_to_q)

	for j := len(p_to_p) - 1; j >= 0; j-- {
		for l := len(p_to_q) - 1; l >= 0; l-- {
			if p_to_p[j] == p_to_q[l] {
				return p_to_p[j]
			}
		}
	}

	return p
}
