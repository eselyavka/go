package solutions

import "sort"

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	d := make(map[int][]int)

	queue := []tuple{tuple{root, 0}}
	var t tuple

	for len(queue) > 0 {
		t = queue[0]

		if _, ok := d[t.col]; ok {
			d[t.col] = append(d[t.col], t.node.Val)
		} else {
			d[t.col] = []int{t.node.Val}
		}

		if t.node.Left != nil {
			queue = append(queue, tuple{t.node.Left, t.col - 1})
		}

		if t.node.Right != nil {
			queue = append(queue, tuple{t.node.Right, t.col + 1})
		}

		queue = queue[1:]
	}

	keys := make([]int, 0)
	res := make([][]int, len(d))

	for key, _ := range d {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		res[i] = d[keys[i]]
	}
	return res
}
