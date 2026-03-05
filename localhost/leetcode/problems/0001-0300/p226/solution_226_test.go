package p226

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution226(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 9, Left: nil, Right: nil}

	curr := invertTree(&root)

	q := make([]*util.TreeNode, 0)
	q = append(q, curr)
	var node *util.TreeNode
	actual := make([]int, 0)

	for len(q) > 0 {
		node = q[0]
		actual = append(actual, node.Val)
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	assert.Equal([]int{4, 7, 2, 9, 6, 3, 1}, actual, "Solution226")
}
