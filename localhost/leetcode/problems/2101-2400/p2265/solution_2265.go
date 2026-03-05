package p2265

import "localhost/leetcode/util"

func subtree(node *util.TreeNode) util.Tuple4 {
	if node == nil {
		return util.Tuple4{Count: 0, Sum: 0}
	}

	t_left := subtree(node.Left)
	t_right := subtree(node.Right)

	c := t_left.Count + t_right.Count
	s := t_left.Sum + t_right.Sum

	return util.Tuple4{Count: 1 + c, Sum: node.Val + s}
}

func averageOfSubtree(root *util.TreeNode) int {
	q := make([]*util.TreeNode, 0)
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

		if node.Val == t_s_c.Sum/t_s_c.Count {
			ans++
		}
	}

	return ans
}
