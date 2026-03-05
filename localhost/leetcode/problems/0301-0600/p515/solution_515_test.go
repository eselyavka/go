package p515

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

func TestSolution515(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 9, Left: nil, Right: nil}

	actual := largestValues(&root)
	assert.Equal(actual, []int{1, 3, 9}, "Solution515")
}
