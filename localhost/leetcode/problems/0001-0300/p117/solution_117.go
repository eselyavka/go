package p117

import "localhost/leetcode/util"

func connect(root *util.NodeNext) *util.NodeNext {
	if root == nil {
		return root
	}

	q := make([]*util.NodeNext, 0)
	q = append(q, root)

	for len(q) > 0 {
		cnt := len(q)
		arr := make([]*util.NodeNext, 0)

		for cnt > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			arr = append(arr, node)
			cnt--
		}
		for i := 1; i < len(arr); i++ {
			arr[i-1].Next = arr[i]
		}
	}

	return root
}
