package p2476

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

func TestSolution2476(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 13, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right.Left = &util.TreeNode{Val: 14, Left: nil, Right: nil}

	actual := closestNodes(&root, []int{2, 5, 16})

	assert.Equal([][]int{{2, 2}, {4, 6}, {15, -1}}, actual, "Solution2476")
}
