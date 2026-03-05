package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution109(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{-10, -3, 0, 5, 9})
	root := sortedListToBST(l)
	q := []*TreeNode{root}
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
