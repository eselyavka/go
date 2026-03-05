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

func TestSolution110(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}

	actual := isBalanced(&root)
	assert.True(actual, "Solution110")
}
