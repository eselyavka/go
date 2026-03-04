package solutions

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	ans := make([]int, 0)

	for len(q) > 0 {
		rowMax := MinInt
		it := len(q)
		for it > 0 {
			node := q[0]
			q = q[1:]
			rowMax = MaxInts([]int{rowMax, node.Val})
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			it--
		}
		ans = append(ans, rowMax)
	}

	return ans
}
