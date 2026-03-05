package p102

import "github.com/eseliavka/go/util"

func levelOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := make([]*util.TreeNode, 0)
	q = append(q, root)

	ans := make([][]int, 0)
	var acc []int

	for len(q) > 0 {
		n := len(q)
		acc = make([]int, 0)
		for n > 0 {
			node := q[0]
			acc = append(acc, node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
			n--
		}
		ans = append(ans, acc)
	}
	return ans
}
