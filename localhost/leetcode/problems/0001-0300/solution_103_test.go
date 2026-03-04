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

func TestSolution103(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	actual := zigzagLevelOrder(&root)

	assert.Equal([][]int{{3}, {20, 9}, {15, 7}}, actual, "Solution103")
}
