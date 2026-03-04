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

func TestSolution1372(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 1, Left: nil, Right: nil}

	actual := longestZigZag(&root)

	assert.Equal(4, actual, "Solution1372")
}
