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

func TestSolution235(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}

	actual := lowestCommonAncestorTreeNode(&root, root.Left, root.Right)

	assert.Equal(6, actual.Val, "Solution235")
}
