package p2415

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

func TestSolution2415(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 8, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 13, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 21, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 34, Left: nil, Right: nil}

	reverseOddLevels(&root)
	queue := make([]*util.TreeNode, 0)
	queue = append(queue, &root)
	actual := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		actual = append(actual, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	assert.Equal(actual, []int{2, 5, 3, 8, 13, 21, 34}, "Solution769")
}
