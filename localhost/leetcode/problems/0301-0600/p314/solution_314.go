package p314

import (
	"localhost/leetcode/util"
	"sort"
)

func verticalOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	d := make(map[int][]int)

	queue := []util.Tuple{{Node: root, Col: 0}}
	var t util.Tuple

	for len(queue) > 0 {
		t = queue[0]

		if _, ok := d[t.Col]; ok {
			d[t.Col] = append(d[t.Col], t.Node.Val)
		} else {
			d[t.Col] = []int{t.Node.Val}
		}

		if t.Node.Left != nil {
			queue = append(queue, util.Tuple{Node: t.Node.Left, Col: t.Col - 1})
		}

		if t.Node.Right != nil {
			queue = append(queue, util.Tuple{Node: t.Node.Right, Col: t.Col + 1})
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
