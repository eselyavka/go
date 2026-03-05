package p515

import "github.com/eseliavka/go/util"

func largestValues(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := make([]*util.TreeNode, 0)
	q = append(q, root)

	ans := make([]int, 0)

	for len(q) > 0 {
		rowMax := util.MinInt
		it := len(q)
		for it > 0 {
			node := q[0]
			q = q[1:]
			rowMax = util.MaxInts([]int{rowMax, node.Val})
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
