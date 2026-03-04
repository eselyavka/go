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

func TestSolution2476(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 13, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right.Left = &TreeNode{Val: 14, Left: nil, Right: nil}

	actual := closestNodes(&root, []int{2, 5, 16})

	assert.Equal([][]int{{2, 2}, {4, 6}, {15, -1}}, actual, "Solution2476")
}
