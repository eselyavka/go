package solutions

func dfs_235(root, node *TreeNode, path *[]*TreeNode) bool {
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

func lowestCommonAncestorTreeNode(root, p, q *TreeNode) *TreeNode {
	p_to_p := make([]*TreeNode, 0)
	p_to_q := make([]*TreeNode, 0)
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
