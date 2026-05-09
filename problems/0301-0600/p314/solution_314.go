package p314

import (
	"github.com/eseliavka/go/util"
	"sort"
)

type columnNode struct {
	node *util.TreeNode
	col  int
}

func verticalOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	d := make(map[int][]int)

	queue := []columnNode{{node: root, col: 0}}
	var t columnNode

	for len(queue) > 0 {
		t = queue[0]

		if _, ok := d[t.col]; ok {
			d[t.col] = append(d[t.col], t.node.Val)
		} else {
			d[t.col] = []int{t.node.Val}
		}

		if t.node.Left != nil {
			queue = append(queue, columnNode{node: t.node.Left, col: t.col - 1})
		}

		if t.node.Right != nil {
			queue = append(queue, columnNode{node: t.node.Right, col: t.col + 1})
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
