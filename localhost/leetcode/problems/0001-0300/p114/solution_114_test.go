package p114

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

func TestSolution114(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 6, Left: nil, Right: nil}

	flatten(&root)

	res := []int{root.Val}
	node := root.Right
	for node != nil {
		res = append(res, node.Val)
		node = node.Right
	}

	assert.Equal(res, []int{1, 2, 3, 4, 5, 6}, "Solution114")
}
