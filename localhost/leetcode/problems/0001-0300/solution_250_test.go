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

func TestSolution250(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	actual := countUnivalSubtrees(&root)

	assert.Equal(4, actual, "Solution250")
}
