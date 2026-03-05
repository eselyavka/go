package p314

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

func TestSolution314(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 0, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 7, Left: nil, Right: nil}

	assert.Equal(verticalOrder(&root), [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}, "Solution314")
}
