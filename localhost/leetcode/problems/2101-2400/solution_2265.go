package solutions

func subtree(node *TreeNode) tuple4 {
	if node == nil {
		return tuple4{count: 0, sum: 0}
	}

	t_left := subtree(node.Left)
	t_right := subtree(node.Right)

	c := t_left.count + t_right.count
	s := t_left.sum + t_right.sum

	return tuple4{count: 1 + c, sum: node.Val + s}
}

func averageOfSubtree(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	ans := 0

	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		t_s_c := subtree(node)

		if node.Val == t_s_c.sum/t_s_c.count {
			ans++
		}
	}

	return ans
}
