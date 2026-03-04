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

func TestSolution1457(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Right.Right = &TreeNode{1, nil, nil}
	assert.Equal(pseudoPalindromicPaths(&root), 2, "Solution1457")
}
