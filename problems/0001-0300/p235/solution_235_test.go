package p235

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

func TestSolution235(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 9, Left: nil, Right: nil}

	actual := lowestCommonAncestorTreeNode(&root, root.Left, root.Right)

	assert.Equal(6, actual.Val, "Solution235")
}
