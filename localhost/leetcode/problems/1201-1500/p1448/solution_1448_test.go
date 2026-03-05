package p1448

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

func TestSolution1448(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}

	actual := goodNodes(&root)
	assert.Equal(4, actual, "Solution1448")
}
