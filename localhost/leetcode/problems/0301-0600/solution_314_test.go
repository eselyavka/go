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

func TestSolution314(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	assert.Equal(verticalOrder(&root), [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}, "Solution314")
}
