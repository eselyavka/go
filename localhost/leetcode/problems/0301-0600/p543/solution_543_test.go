package p543

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

func TestSolution543(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}

	actual := diameterOfBinaryTree(&root)
	assert.Equal(3, actual, "Solution543")
}
