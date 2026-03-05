package p1457

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

func TestSolution1457(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	assert.Equal(pseudoPalindromicPaths(&root), 2, "Solution1457")
}
