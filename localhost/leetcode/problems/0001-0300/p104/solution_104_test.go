package p104

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

func TestSolution104(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 7, Left: nil, Right: nil}

	actual := maxDepth(&root)

	assert.Equal(3, actual, "Solution104")

}
