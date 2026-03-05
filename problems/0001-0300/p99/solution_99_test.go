package p99

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution99(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 2, Left: nil, Right: nil}

	recoverTree(&root)

	res := make([]int, 0)
	q := make([]*util.TreeNode, 0)
	q = append(q, &root)
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		res = append(res, node.Val)

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	assert.Equal([]int{3, 1, 2}, res, "Solution99")
}
