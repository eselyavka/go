package p426

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

func TestSolution426(t *testing.T) {
	assert := assert.New(t)
	root := util.Node{Val: 4, Left: nil, Right: nil, Parent: nil}
	root.Left = &util.Node{Val: 2, Left: nil, Right: nil, Parent: nil}
	root.Left.Left = &util.Node{Val: 1, Left: nil, Right: nil, Parent: nil}
	root.Left.Right = &util.Node{Val: 3, Left: nil, Right: nil, Parent: nil}
	root.Right = &util.Node{Val: 5, Left: nil, Right: nil, Parent: nil}

	res := treeToDoublyList(&root)
	var head *util.Node = nil

	actual := make([]int, 0)

	for head == nil || res.Val != head.Val {
		if head == nil {
			head = res
		}
		actual = append(actual, res.Val)
		res = res.Right
	}

	assert.Equal(actual, []int{1, 2, 3, 4, 5}, "Solution426")
}
