package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution226(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}

	curr := invertTree(&root)

	q := make([]*TreeNode, 0)
	q = append(q, curr)
	var node *TreeNode
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
