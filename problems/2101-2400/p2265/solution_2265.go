package p2265

import "github.com/eseliavka/go/util"

type subtreeStats struct {
	count int
	sum   int
}

func subtree(node *util.TreeNode) subtreeStats {
	if node == nil {
		return subtreeStats{count: 0, sum: 0}
	}

	t_left := subtree(node.Left)
	t_right := subtree(node.Right)

	c := t_left.count + t_right.count
	s := t_left.sum + t_right.sum

	return subtreeStats{count: 1 + c, sum: node.Val + s}
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

		if node.Val == t_s_c.sum/t_s_c.count {
			ans++
		}
	}

	return ans
}
