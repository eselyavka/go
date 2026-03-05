package p1372

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

func TestSolution1372(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}

	actual := longestZigZag(&root)

	assert.Equal(4, actual, "Solution1372")
}
