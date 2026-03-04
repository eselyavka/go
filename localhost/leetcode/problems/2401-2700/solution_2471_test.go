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

func TestSolution2471(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right.Left = &TreeNode{Val: 10, Left: nil, Right: nil}

	actual := minimumOperations_2471(&root)
	assert.Equal(actual, 3, "Solution2471")
}
