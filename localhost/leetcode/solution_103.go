package solutions

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	ans := make([][]int, 0)

	dir := 0

	for len(q) > 0 {
		i := len(q)
		buf := make([]int, 0)
		for i > 0 {
			node := q[0]
			q = q[1:]
			if dir%2 == 1 {
				buf = append([]int{node.Val}, buf...)
			} else {
				buf = append(buf, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			i--
		}
		dir++
		ans = append(ans, buf)
	}

	return ans
}
