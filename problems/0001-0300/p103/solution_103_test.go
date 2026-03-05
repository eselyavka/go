package p103

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

func TestSolution103(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 7, Left: nil, Right: nil}

	actual := zigzagLevelOrder(&root)

	assert.Equal([][]int{{3}, {20, 9}, {15, 7}}, actual, "Solution103")
}
