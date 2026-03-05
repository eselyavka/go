package p109

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

func TestSolution109(t *testing.T) {
	assert := assert.New(t)
	l := util.InitLinkedList([]int{-10, -3, 0, 5, 9})
	root := sortedListToBST(l)
	q := []*util.TreeNode{root}
	actual := make([]int, 0)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		actual = append(actual, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	assert.Equal([]int{0, -3, 9, -10, 5}, actual, "Solution109")
}
